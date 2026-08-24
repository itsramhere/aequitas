import json
import glob
import os
import matplotlib.pyplot as plt

def load_data(results_dir):
    data = []
    filepaths = glob.glob(os.path.join(results_dir, "*.json"))
    
    for filepath in filepaths:
        with open(filepath, 'r') as f:
            try:
                cell = json.load(f)
                data.append(cell)
            except json.JSONDecodeError as e:
                print(f"Error reading {filepath}: {e}")
    return data

def extract_series(data, fixed_key, fixed_val, x_key, y_func):
    strategies = ["OCC", "PESSIMISTIC", "SSI", "ADAPTIVE"]
    series_data = {strat: {"x": [], "y": []} for strat in strategies}
    
    for cell in data:
        if cell.get(fixed_key) == fixed_val:
            strat = cell.get("strategy")
            if strat in strategies:
                x_val = cell.get(x_key)
                y_val = y_func(cell)
                
                series_data[strat]["x"].append(x_val)
                series_data[strat]["y"].append(y_val)
                
    # Sort the coordinates by x-value to ensure lines plot correctly
    for strat in strategies:
        sorted_pairs = sorted(zip(series_data[strat]["x"], series_data[strat]["y"]))
        if sorted_pairs:
            series_data[strat]["x"], series_data[strat]["y"] = zip(*sorted_pairs)
            
    return series_data

def plot_graph(series_data, title, xlabel, ylabel, filename, y_log_scale=False):
    plt.figure(figsize=(10, 6))
    
    colors = {"OCC": "blue", "PESSIMISTIC": "red", "SSI": "green", "ADAPTIVE": "purple"}
    markers = {"OCC": "o", "PESSIMISTIC": "s", "SSI": "^", "ADAPTIVE": "d"}
    
    for strat, data in series_data.items():
        if data["x"]:
            plt.plot(data["x"], data["y"], label=strat, color=colors[strat], 
                     marker=markers[strat], linewidth=2, markersize=8)
            
    plt.title(title, fontsize=14, fontweight="bold")
    plt.xlabel(xlabel, fontsize=12)
    plt.ylabel(ylabel, fontsize=12)
    plt.grid(True, linestyle="--", alpha=0.7)
    plt.legend(fontsize=12)
    
    if y_log_scale:
        plt.yscale("log")
        
    plt.tight_layout()
    plt.savefig(filename, dpi=300)
    print(f"Saved plot to {filename}")
    plt.close()

def main():
    results_dir = os.path.join("results", "set1")
    if not os.path.exists(results_dir):
        print(f"Directory {results_dir} not found.")
        return
        
    data = load_data(results_dir)
    print(f"Loaded {len(data)} JSON reports.")
    
    plots_dir = "plots"
    if not os.path.exists(plots_dir):
        os.makedirs(plots_dir)

    # 1. Throughput vs. Skew (Fixed Concurrency = 50)
    skew_series = extract_series(
        data, 
        fixed_key="concurrency", fixed_val=50, 
        x_key="skew_theta", 
        y_func=lambda c: c.get("throughput_tps", 0)
    )
    plot_graph(skew_series, "Throughput vs. Zipfian Skew (Concurrency = 50)", 
               "Zipfian Skew (\u03b8)", "Throughput (TPS)", 
               os.path.join(plots_dir, "throughput_vs_skew.png"))

    # 2. Throughput vs. Concurrency (Fixed Skew = 0.8)
    conc_series = extract_series(
        data, 
        fixed_key="skew_theta", fixed_val=0.8, 
        x_key="concurrency", 
        y_func=lambda c: c.get("throughput_tps", 0)
    )
    plot_graph(conc_series, "Throughput vs. Concurrency (Skew \u03b8 = 0.8)", 
               "Concurrent Clients", "Throughput (TPS)", 
               os.path.join(plots_dir, "throughput_vs_concurrency.png"))

    # 3. Client Failure Rate vs. Skew (Fixed Concurrency = 50)
    failure_series = extract_series(
        data, 
        fixed_key="concurrency", fixed_val=50, 
        x_key="skew_theta", 
        y_func=lambda c: c.get("client_failure_rate_pct", 0)
    )
    plot_graph(failure_series, "Client Failure Rate vs. Zipfian Skew (Concurrency = 50)", 
               "Zipfian Skew (\u03b8)", "Failure Rate (%)", 
               os.path.join(plots_dir, "failure_rate_vs_skew.png"))

    # 4. P99 CC Wait Latency vs. Concurrency (Fixed Skew = 0.8)
    # Convert nanoseconds to milliseconds by dividing by 1,000,000
    latency_series = extract_series(
        data, 
        fixed_key="skew_theta", fixed_val=0.8, 
        x_key="concurrency", 
        y_func=lambda c: c.get("cc_wait_latency", {}).get("p99", 0) / 1_000_000.0
    )
    plot_graph(latency_series, "P99 CC Wait Latency vs. Concurrency (Skew \u03b8 = 0.8)", 
               "Concurrent Clients", "P99 Wait Latency (ms)", 
               os.path.join(plots_dir, "cc_wait_latency_vs_concurrency.png"),
               y_log_scale=True) # Log scale helps visualize the massive queue spike

if __name__ == "__main__":
    main()