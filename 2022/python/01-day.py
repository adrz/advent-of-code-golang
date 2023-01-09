from utils import download_input

if __name__ == "__main__":
    data = download_input(1)
    data_list = data.split("\n\n")
    data_list[-1] = data_list[-1][:-1]

    data_sum = [sum([int(x) for x in group.split("\n")]) for group in data_list]
    print(max(data_sum))
    print(sum(sorted(data_sum)[-3:]))
