import os
import glob
import json
import numpy as np
import matplotlib.pyplot as plt

def load_benchmark_data():
    search_dirs = [
        os.path.join("results", "set1_ci"),
        os.path.join("results", "set1")
    ]
    
    data_files = []
    for d in search_dirs:
        if os.path.exists(d):
            files = glob.glob(os.path.join(d, "*.json"))
            if files:
                data_files.extend(files)
                print(f"Found {len(files)} JSON report files in {d}")

    data_cells = []
    for filepath in data_files:
        try:
            with open(filepath, "r", encoding="utf-8") as f:
                cell = json.load(f)
                data_cells.append(cell)
        except Exception as e:
            print(f"Error loading {filepath}: {e}")

    return data_cells

def aggregate_matrix(data_cells, strategy):
    skews = [0.0, 0.4, 0.6, 0.8, 1.0, 1.2]
    concs = [10, 50, 100, 250]

    grid = np.zeros((len(concs), len(skews)))
    counts = np.zeros((len(concs), len(skews)))

    for cell in data_cells:
        if cell.get("strategy") == strategy:
            s_val = float(cell.get("skew_theta", 0.0))
            c_val = int(cell.get("concurrency", 0))
            tps = float(cell.get("throughput_tps", 0.0))

            if s_val in skews and c_val in concs:
                s_idx = skews.index(s_val)
                c_idx = concs.index(c_val)
                grid[c_idx, s_idx] += tps
                counts[c_idx, s_idx] += 1

    with np.errstate(divide='ignore', invalid='ignore'):
        mean_grid = np.where(counts > 0, grid / counts, 0.0)

    return mean_grid, skews, concs

def main():
    data_cells = load_benchmark_data()
    if not data_cells:
        print("No benchmark JSON data found in results directory.")
        return

    strategies = ["OCC", "PESSIMISTIC", "SSI", "ADAPTIVE"]
    plots_dir = "plots"
    if not os.path.exists(plots_dir):
        os.makedirs(plots_dir)

    fig, axes = plt.subplots(2, 2, figsize=(14, 10))
    axes = axes.flatten()

    for idx, strat in enumerate(strategies):
        ax = axes[idx]
        grid, skews, concs = aggregate_matrix(data_cells, strat)

        im = ax.imshow(grid, cmap="YlGnBu", aspect="auto", origin="lower")
        
        ax.set_xticks(range(len(skews)))
        ax.set_xticklabels([f"{s:.1f}" for s in skews], fontsize=10)
        ax.set_yticks(range(len(concs)))
        ax.set_yticklabels(concs, fontsize=10)

        ax.set_xlabel("Zipfian Skew (Theta)", fontsize=11, fontweight="bold")
        ax.set_ylabel("Concurrent Clients", fontsize=11, fontweight="bold")
        ax.set_title(f"Strategy: {strat}", fontsize=13, fontweight="bold")

        # Annotate text values inside heatmap cells
        for i in range(len(concs)):
            for j in range(len(skews)):
                val = grid[i, j]
                if val > 0:
                    text_color = "white" if val > np.max(grid) * 0.65 else "black"
                    ax.text(j, i, f"{val:.0f}", ha="center", va="center", color=text_color, fontsize=9, fontweight="bold")

        cbar = fig.colorbar(im, ax=ax, shrink=0.8)
        cbar.set_label("Throughput (TPS)", fontsize=10)

    plt.suptitle("Contention Heatmap: Mean Throughput (TPS) across Skew and Concurrency", fontsize=16, fontweight="bold", y=0.98)
    plt.tight_layout(rect=[0, 0, 1, 0.95])

    out_file = os.path.join(plots_dir, "contention_heatmap.png")
    plt.savefig(out_file, dpi=300)
    print(f"Saved contention heatmap plot to {out_file}")
    plt.close()

if __name__ == "__main__":
    main()
