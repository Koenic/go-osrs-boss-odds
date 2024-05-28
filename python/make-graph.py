import json
import os
import matplotlib.pyplot as plt
from scipy.interpolate import interp1d
import numpy as np
from cycler import cycler
from math import atan2,degrees

#Label line with line2D label data
def labelLine(line,x,label=None,align=True,**kwargs):

    ax = line.axes
    xdata = line.get_xdata()
    ydata = line.get_ydata()

    if (x < xdata[0]) or (x > xdata[-1]):
        return

    #Find corresponding y co-ordinate and angle of the line
    ip = 1
    for i in range(len(xdata)):
        if x < xdata[i]:
            ip = i
            break

    y = ydata[ip-1] + (ydata[ip]-ydata[ip-1])*(x-xdata[ip-1])/(xdata[ip]-xdata[ip-1])

    if not label:
        label = line.get_label()

    if align:
        #Compute the slope
        dx = xdata[ip] - xdata[ip-1]
        dy = ydata[ip] - ydata[ip-1]
        ang = degrees(atan2(dy,dx))

        #Transform to screen co-ordinates
        pt = np.array([x,y]).reshape((1,2))
        trans_angle = ax.transData.transform_angles(np.array((ang,)),pt)[0]

    else:
        trans_angle = 0

    #Set a bunch of keyword arguments
    if 'color' not in kwargs:
        kwargs['color'] = line.get_color()

    if ('horizontalalignment' not in kwargs) and ('ha' not in kwargs):
        kwargs['ha'] = 'center'

    if ('verticalalignment' not in kwargs) and ('va' not in kwargs):
        kwargs['va'] = 'center'

    if 'backgroundcolor' not in kwargs:
        kwargs['backgroundcolor'] = ax.get_facecolor()

    if 'clip_on' not in kwargs:
        kwargs['clip_on'] = True

    if 'zorder' not in kwargs:
        kwargs['zorder'] = 2.5

    ax.text(x,y,label,rotation=trans_angle,**kwargs)

def labelLines(lines,align=True,xvals=None,**kwargs):

    ax = lines[0].axes
    labLines = []
    labels = []

    #Take only the lines which have labels other than the default ones
    for line in lines:
        label = line.get_label()
        if "_line" not in label:
            labLines.append(line)
            labels.append(label)

    if xvals is None:
        xmin,xmax = ax.get_xlim()
        xvals = np.linspace(xmin,xmax,len(labLines)+2)[1:-1]

    for line,x,label in zip(labLines,xvals,labels):
        labelLine(line,x,label,align,**kwargs)

def plot_graph(json_file):
    # Load JSON data
    with open(json_file, 'r') as file:
        data = json.load(file)

    cdf_cutoff = 0.99

    monster_name = data["Monster_name"]
    print(monster_name)
    description = data["Description"]
    kc_name = data["Kc_name"]

    fig, ax1 = plt.subplots(figsize=(12, 7))
    ax2 = ax1.twinx()
    ax1.set_xlabel(kc_name)
    ax1.grid(False)
    ax1.grid(True, 'major', axis='x')
    ax2.grid(True)


    blues = plt.get_cmap("Blues", len(data["Variation"]) + 3)
    reds = plt.get_cmap("Reds", len(data["Variation"]) + 3)
    greys = plt.get_cmap("Greys", len(data["Variation"]))

    
    blue_col = 'dodgerblue'
    red_col = 'red'
    ax1.set_ylabel(f'Odds of completion at specific {kc_name}', color=red_col)
    ax2.set_ylabel(f'Odds of completion before a {kc_name}', color=blue_col)
    ax1.tick_params(axis='y', labelcolor=red_col)
    ax2.tick_params(axis='y', labelcolor=blue_col)
    
    ax1yLim = 0
    ax2yLim = 0
    idx_cutoff = 0
    idx_max_cutoff = 1000000

    # interpolate data
    interpolated_variation = []
    for index, variation in enumerate(data["Variation"]):
        variation_name = variation["VariationName"]
        cdf_values = variation["Cdf"]
        cdf_x_vals = variation["X_vals"]

        # Interpolate the CDF values at the new x-values
        interpolation_function = interp1d(cdf_x_vals, cdf_values, kind='cubic', fill_value="extrapolate")
        cdf_x_vals_new = np.arange(0, max(cdf_x_vals))
        cdf_values_new = interpolation_function(cdf_x_vals_new)

        # Calculate PDF values from CDF values
        pdf_values = np.diff(cdf_values_new, prepend=cdf_values_new[0])
        pdf_x_values = cdf_x_vals_new[1:]
        pdf_values = pdf_values[1:]

        # Find mean/mode/median
        median_index = np.abs(np.array(cdf_values_new) - 0.5).argmin()
        mode_index = np.argmax(pdf_values)
        mean_index = np.sum(pdf_x_values * pdf_values) / np.sum(pdf_values)

        # Find the index where CDF reaches the cutoff value
        
        if(idx_max_cutoff > np.argmax(cdf_x_vals_new)):
            idx_max_cutoff = np.argmax(cdf_x_vals_new)
        if(cdf_x_vals_new[np.argmax(np.array(cdf_values_new) >= cdf_cutoff)] > idx_cutoff):
            idx_cutoff = cdf_x_vals_new[np.argmax(np.array(cdf_values_new) >= cdf_cutoff)]
        

        if(pdf_values[np.argmax(pdf_values[:idx_cutoff])] > ax1yLim):
            ax1yLim = pdf_values[np.argmax(pdf_values[:idx_cutoff])]
        if(cdf_values_new[np.argmax(cdf_values_new[:idx_cutoff])] > ax2yLim):
            ax2yLim = cdf_values_new[np.argmax(cdf_values_new[:idx_cutoff])]

        variation = {
            "pdf_values": pdf_values,
            "pdf_x_values": pdf_x_values,
            "cdf_values": cdf_values_new,
            "cdf_x_values": cdf_x_vals_new,
            "mean_index": mean_index,
            "mode_index": mode_index,
            "median_index": median_index,
            "idx_cutoff": idx_cutoff,
            "variation_name": variation_name
        }
        interpolated_variation.append(variation)

    if(idx_max_cutoff < idx_cutoff):
        idx_cutoff = idx_max_cutoff
    
    ax1.set_xlim(right=idx_cutoff)
    ax1.set_xlim(left=-idx_cutoff/100)
    ax2.set_ylim(top=ax2yLim * 1.05, bottom=-ax2yLim * 0.02)
    ax1.set_ylim(top=ax1yLim * 1.05, bottom=-ax1yLim * 0.02)

    for index, variation in enumerate(interpolated_variation):
        pdf_values = variation["pdf_values"] 
        pdf_x_values = variation["pdf_x_values"] 
        cdf_values = variation["cdf_values"] 
        cdf_x_values = variation["cdf_x_values"] 
        mean_index = variation["mean_index"] 
        mode_index = variation["mode_index"] 
        median_index = variation["median_index"]
        variation_name = variation["variation_name"]

        blue_col = blues(index + 2)
        red_col = reds(index + 2)
        grey_col = greys(index + 1)


        if(len(data["Variation"]) == 1):
            ax2.plot(cdf_x_values[median_index], cdf_values[median_index], marker='o', markersize=8, color=blue_col, label=f'Half of players complete at or before {cdf_x_values[median_index]} {kc_name} (median)')
            ax1.plot(pdf_x_values[mode_index], pdf_values[mode_index], marker='o', markersize=8, color=red_col, label='Most ~1/{:.0f} players complete at {:.0f} {} (mode)'.format(1/pdf_values[mode_index], pdf_x_values[mode_index], kc_name ))
            ax1.axvline(x=mean_index, color=grey_col, linestyle='--', label='Average {} at completion: ~{:.2f} (mean)'.format(kc_name, mean_index))
            ax2.plot(cdf_x_values, cdf_values, linestyle='-', color=blue_col)
            ax1.plot(pdf_x_values, pdf_values, linestyle='-', color=red_col)
        else:
            ax2.plot(cdf_x_values, cdf_values, linestyle='-', color=blue_col, label=f'{variation_name}')
            ax1.plot(pdf_x_values, pdf_values, linestyle='-', color=red_col, label=f'{variation_name}')


    # Title and legend
    plt.title(f'{monster_name}')
    if(len(data["Variation"]) > 1):
        labelLines(ax2.get_lines(),zorder=2.5)
        labelLines(ax1.get_lines(),zorder=2.5)
    else:
        fig.legend(loc='upper right')
    fig.tight_layout()

    lines = 1
    if(len(description) > 110):
        descParts = description.split(', ')
        description = ''
        partlen = 0
        for ind, descPart in enumerate(descParts):
            partlen += len(descPart)
            description += descPart + ', '
            if(partlen > 110 and ind < len(descParts) - 1):
                lines += 1
                partlen = 0
                description += '\n'
        description = description[:-2]
    

    plt.figtext(0.5, 0.02, description, ha='center', fontsize=10)
    
    if(len(data["Variation"]) == 1):
        plt.subplots_adjust(bottom=0.08 + lines * 0.03, top=0.85)
    else:
        plt.subplots_adjust(bottom=0.08 + lines * 0.03)
        
    # plt.subplots_adjust(bottom=0.15, top=0.85)
    plt.autoscale(False)

    # Save the plot as image files
    os.makedirs('plots', exist_ok=True)
    plt.savefig(os.path.join('plots', os.path.basename(json_file).split('.')[0] + ".png"))

    # Ensure the plot is saved and the script exits cleanly
    plt.close()

data_folder = 'data'
for file_name in os.listdir(data_folder):
    if file_name.endswith(".json"):
        json_file = os.path.join(data_folder, file_name)
        plot_graph(json_file)