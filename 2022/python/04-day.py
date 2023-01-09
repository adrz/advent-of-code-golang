from utils import download_input

if __name__ == "__main__":
    data = download_input(4)

    data_list = data.splitlines()
    completely_overlapped_pairs = 0
    non_overlapping_pairs = 0

    for pair in data_list:
        [elf_one, elf_two] = pair.split(",")

        [start1, end1] = [int(x) for x in elf_one.split("-")]
        [start2, end2] = [int(x) for x in elf_two.split("-")]

        # part 1
        if (start1 <= start2 and end2 <= end1) or (start2 <= start1 and end1 <= end2):
            completely_overlapped_pairs += 1

        # part 2
        if (start1 < start2 and end1 < start2) or (start2 < start1 and end2 < start1):
            non_overlapping_pairs += 1

    # part 1
    print(completely_overlapped_pairs)

    # part 2
    overlapped_pairs = len(data_list) - non_overlapping_pairs
    print(overlapped_pairs)
