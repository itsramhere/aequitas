import json
import glob
import os
import math
import sys
from collections import defaultdict

# Two-sided 95% Student-t critical values by degrees of freedom (df = n - 1).
# With n = 3..5 repetitions the normal z = 1.96 understates the half-width by
# roughly 2x; the t distribution is the correct small-sample interval.
T_CRIT_95 = {
    1: 12.706, 2: 4.303, 3: 3.182, 4: 2.776, 5: 2.571,
    6: 2.447, 7: 2.365, 8: 2.306, 9: 2.262, 10: 2.228,
    15: 2.131, 20: 2.086, 30: 2.042,
}

def t_critical(df):
    if df in T_CRIT_95:
        return T_CRIT_95[df]
    if df > 30:
        return 1.96  # normal approximation is fine for large samples
    # Interpolate between the two nearest tabulated df values
    keys = sorted(T_CRIT_95)
    lo = max(k for k in keys if k <= df)
    hi = min(k for k in keys if k >= df)
    if lo == hi:
        return T_CRIT_95[lo]
    frac = (df - lo) / (hi - lo)
    return T_CRIT_95[lo] + frac * (T_CRIT_95[hi] - T_CRIT_95[lo])

def calculate_stats(data_list):
    n = len(data_list)
    if n == 0:
        return 0, 0
    mean = sum(data_list) / n
    if n == 1:
        return mean, 0
    variance = sum((x - mean) ** 2 for x in data_list) / (n - 1)
    std_dev = math.sqrt(variance)
    ci_95 = t_critical(n - 1) * (std_dev / math.sqrt(n))
    return mean, ci_95

def main():
    # Exactly one result directory per invocation, so homogeneous repetition
    # sets are never pooled with single-run or different-repetition-count data.
    if len(sys.argv) > 1:
        search_dirs = [sys.argv[1]]
    else:
        search_dirs = [os.path.join("results", "set1_ci")]

    filepaths = []
    for d in search_dirs:
        if os.path.exists(d):
            files = glob.glob(os.path.join(d, "*.json"))
            if files:
                filepaths.extend(files)
                print(f"Found {len(files)} JSON report files in {d}")

    if not filepaths:
        print("No benchmark result JSON files found.")
        return

    # Structure: [skew][strategy] = {"tps": [], "retries": [], "failures_pct": []}
    data = defaultdict(lambda: defaultdict(lambda: {"tps": [], "retries": [], "failures_pct": []}))

    for filepath in filepaths:
        with open(filepath, 'r') as f:
            try:
                cell = json.load(f)
                strat = cell.get("strategy")
                skew = cell.get("skew_theta")
                if strat and skew is not None:
                    data[skew][strat]["tps"].append(cell.get("throughput_tps", 0))
                    data[skew][strat]["retries"].append(cell.get("internal_retries", 0))
                    data[skew][strat]["failures_pct"].append(cell.get("client_failure_rate_pct", 0))
            except json.JSONDecodeError:
                pass

    print("=== TABLE I: EMPIRICAL COMPARISON WITH 95% CIs (Student-t) ===")
    print(f"{'Strategy':<15} | {'Skew':<5} | {'n':<3} | {'Throughput (TPS)':<20} | {'Internal Retries':<20} | {'Client Failures (%)'}")
    print("-" * 95)

    strat_order = ["OCC", "PESSIMISTIC", "ADAPTIVE", "SSI"]

    for skew in sorted(data.keys()):
        for strat in strat_order:
            if strat in data[skew]:
                metrics = data[skew][strat]
                tps_mean, tps_ci = calculate_stats(metrics["tps"])
                retries_mean, retries_ci = calculate_stats(metrics["retries"])
                fail_mean, fail_ci = calculate_stats(metrics["failures_pct"])

                tps_str = f"{tps_mean:.1f} \u00B1 {tps_ci:.1f}"
                retries_str = f"{int(retries_mean)} \u00B1 {int(retries_ci)}"
                fail_str = f"{fail_mean:.2f}% \u00B1 {fail_ci:.2f}%"

                print(f"{strat:<15} | {skew:<5.1f} | {len(metrics['tps']):<3} | {tps_str:<20} | {retries_str:<20} | {fail_str}")

if __name__ == "__main__":
    main()
