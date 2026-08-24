import os
import glob
import json
import matplotlib.pyplot as plt

def find_target_json(strategy, skew=0.8, conc=50):
    search_dirs = [
        os.path.join("results", "set1_ci"),
        os.path.join("results", "set1"),
        os.path.join("results", "set2")
    ]

    for d in search_dirs:
        if os.path.exists(d):
            files = glob.glob(os.path.join(d, "*.json"))
            for f in sorted(files):
                try:
                    with open(f, "r", encoding="utf-8") as fp:
                        data = json.load(fp)
                        if data.get("strategy") == strategy and float(data.get("skew_theta", 0)) == skew and int(data.get("concurrency", 0)) == conc:
                            # ADR-24: search order spans instrument vintages
                            # (set1_ci repetitions vs single-run set1/set2);
                            # always disclose which file a curve came from.
                            print(f"[provenance] {strategy}: {f}")
                            return data
                except Exception:
                    pass
    return None

def extract_timeseries(data, strategy):
    if data and "time_series" in data and len(data["time_series"]) > 0:
        ts_list = data["time_series"]
        seconds = [pt["second"] for pt in ts_list]
        tps = [pt["tps"] for pt in ts_list]
        retries = [pt["retries"] for pt in ts_list]
        return seconds, tps, retries
    else:
        raise SystemExit(
            f"ERROR: no real time_series data found for strategy {strategy}. "
            "Synthetic fallback data was removed so fabricated curves can never "
            "end up in a paper artifact; re-run the benchmark cell first."
        )

def main():
    occ_data = find_target_json("OCC", skew=0.8, conc=50)
    adaptive_data = find_target_json("ADAPTIVE", skew=0.8, conc=50)

    occ_sec, occ_tps, occ_retries = extract_timeseries(occ_data, "OCC")
    adap_sec, adap_tps, adap_retries = extract_timeseries(adaptive_data, "ADAPTIVE")

    plots_dir = "plots"
    if not os.path.exists(plots_dir):
        os.makedirs(plots_dir)

    fig, (ax1, ax2) = plt.subplots(2, 1, figsize=(11, 8), sharex=True)

    # Subplot 1: Throughput (TPS) over time
    ax1.plot(occ_sec, occ_tps, label="OCC", color="blue", linewidth=2, marker="o", markersize=4)
    ax1.plot(adap_sec, adap_tps, label="ADAPTIVE", color="purple", linewidth=2.5, marker="s", markersize=4)
    ax1.set_ylabel("Throughput (TPS)", fontsize=12, fontweight="bold")
    ax1.set_title("Transient Throughput & Retries Stabilization (Skew \u03b8 = 0.8, Concurrency = 50)", fontsize=14, fontweight="bold")
    ax1.grid(True, linestyle="--", alpha=0.7)
    ax1.legend(fontsize=11, loc="upper right")

    # Subplot 2: Retries per second over time
    ax2.plot(occ_sec, occ_retries, label="OCC Retries/sec", color="red", linewidth=2, linestyle="--", marker="^", markersize=4)
    ax2.plot(adap_sec, adap_retries, label="ADAPTIVE Retries/sec", color="green", linewidth=2.5, linestyle="-", marker="d", markersize=4)
    ax2.set_xlabel("Time (Seconds)", fontsize=12, fontweight="bold")
    ax2.set_ylabel("Internal Retries / sec", fontsize=12, fontweight="bold")
    ax2.grid(True, linestyle="--", alpha=0.7)
    ax2.legend(fontsize=11, loc="upper right")

    plt.tight_layout()
    out_file = os.path.join(plots_dir, "transient_stabilization.png")
    plt.savefig(out_file, dpi=300)
    print(f"Saved transient stabilization plot to {out_file}")
    plt.close()

if __name__ == "__main__":
    main()
