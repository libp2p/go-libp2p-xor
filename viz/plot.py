import matplotlib.pyplot as plt
import matplotlib.lines as mlines

from .model import *


def plot(model: LookupModel):
    fig, ax = plt.subplots(1)
    plot_lookup(ax, model)
    plt.show()


def plot_lookup(ax, m: LookupModel):
    # plot vertical grid lines on event stamps
    for e in model.events:
        v_line = mlines.Line2D(
            [m.stamp_to_x(e.stamp_ns), m.stamp_to_x(e.stamp_ns)],
            [0, 1.0],
            color='#d0d0d0')
        ax.add_line(v_line)

    # plot horizontal grid lines on peers that were used
    for u in model.used:
        h_line = mlines.Line2D(
            [m.stamp_to_x(m.start_ns), m.stamp_to_x(m.stop_ns)],
            [m.key_to_y(u), m.key_to_y(u)],
            color='#d0d0d0')
        ax.add_line(h_line)

    # plot state changes
    x, y, s, c = [], [], [], []

    def push(e_, k_, c_):
        ex, ey = m.event_key_xy(e_, k_)
        x.append(ex)
        y.append(ey)
        c.append(c_)
        s.append(1.0)

    for e in model.events:
        for k in e.heard():
            push(e, k, '#80e0c0')
        for k in e.waiting():
            push(e, k, '#80c0e0')
        for k in e.queried():
            push(e, k, '#d0a0e0')
        for k in e.unreachable():
            push(e, k, '#e0a0b0')
        ax.scatter(x, y, s=s, c=c, alpha=0.7, zorder=5, marker='s')

    # plot queries
    for q in model.queries:
        q_line = mlines.Line2D(
            [m.stamp_to_x(q.request.stamp_ns), m.stamp_to_x(q.response.stamp_ns)],
            [m.key_to_y(q.peer), m.key_to_y(q.peer)],
            color=color_for_query_outcome(q))
        ax.add_line(q_line)

    # plot lookup path
    for q in model.find_path:
        q_line = mlines.Line2D(
            [m.stamp_to_x(q.request.stamp_ns), m.stamp_to_x(q.response.stamp_ns)],
            [m.key_to_y(q.peer), m.key_to_y(q.peer)],
            linewidth=3.0,
            color='#e0e0c0')
        ax.add_line(q_line)

    # customize axes
    set_yticks_for_model(ax, m)
    set_yticks_for_model(ax, m)
    style_axis(ax)
    ax.set_title("lookup {}".format(m.id))


def color_for_query_outcome(q):
    if q.outcome == QUERY_SUCCESS:
        return '#50c050'
    elif q.outcome == QUERY_UNREACHABLE:
        return '#c05050'
    else:
        return '#805080'


def set_xticks_for_model(ax, m: LookupModel):
    ax.set_xticks([m.stamp_to_x(e.stamp_ns) for e in m.events])
    ax.set_xticklabels([m.stamp_to_x(e.stamp_ns) for e in m.events])
    ax.set_xlim([m.min_x(), m.max_x()])


def set_yticks_for_model(ax, m: LookupModel):
    ax.set_yticks([m.key_to_y(u) for u in m.used])
    ax.set_yticklabels([m.key_to_y(u) for u in m.used])
    ax.set_ylim([m.min_y(), m.max_y()])


def style_axis(ax):
    ax.grid(zorder=0)
    ax.tick_params(axis='both', which='major', labelsize=6)
    ax.tick_params(axis='both', which='minor', labelsize=5)
