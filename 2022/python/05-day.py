import re

from utils import download_input

if __name__ == "__main__":
    data = download_input(5)

    drawing = data.split("\n\n")

    drawing_of_crates = drawing[0].split("\n")
    rearrangement_procedure = drawing[1].split("\n")[:-1]

    horizontal_crates_drawing = [
        re.findall(r".{1,4}", row) for row in drawing_of_crates[:-1]
    ]

    initial_stacks_of_crates = {}

    # create initial stacks from drawing
    for row in horizontal_crates_drawing:
        for index, crate in enumerate(row):
            crate = crate.strip()
            if crate != "":
                if index + 1 in initial_stacks_of_crates:
                    initial_stacks_of_crates[index + 1].append(crate)
                else:
                    initial_stacks_of_crates[index + 1] = [crate]

    for k, v in initial_stacks_of_crates.items():
        initial_stacks_of_crates[k] = v[::-1]

    for arr in rearrangement_procedure:
        regex = re.search(r"move (\d+) from (\d+) to (\d+)", arr)
        n_to_move, from_stack, to_stack = [int(x) for x in regex.groups()]
        for i in range(n_to_move):
            initial_stacks_of_crates[to_stack].append(
                initial_stacks_of_crates[from_stack].pop()
            )

    my_str = ""
    for i in range(1, 10):
        my_str += initial_stacks_of_crates[i][-1]
