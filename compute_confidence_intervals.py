import json
import glob
import os
import math
from collections import defaultdict

def calculate_stats(data_list):
    n = len(data_list)
    if n == 0:
        return 0, 0
    mean = sum(data_list) / n
    if n == 1:
        return mean, 0
    variance = sum((x - mean) ** 2 for x in data_list) / (n - 1)
    std_dev = math.sqrt(variance)
    ci_95 = 1.96 * (std_dev / math.sqrt(n))
    return mean, ci_95

def main():
    search_dirs = [
        os.path.join("results", "set1_ci"),
        os.path.join("results", "set1_repetitions"),
        os.path.join("results", "set1")
    ]
    
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

    print("=== TABLE I: EMPIRICAL COMPARISON WITH 95% CIs ===")
    print(f"{'Strategy':<15} | {'Skew':<5} | {'Throughput (TPS)':<20} | {'Internal Retries':<20} | {'Client Failures (%)'}")
    print("-" * 85)

    # Sort skews and define strategy print order to match your LaTeX draft
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
                
                print(f"{strat:<15} | {skew:<5.1f} | {tps_str:<20} | {retries_str:<20} | {fail_str}")

if __name__ == "__main__":
    main()