import string

from utils import download_input


def get_v(x):
    return string.ascii_letters.index(x) + 1


def generator(data):
    for start in range(0, len(data), 3):
        yield data[start : start + 3]


if __name__ == "__main__":
    data = download_input(3)
    data_list = data.splitlines()
    total = 0
    for elem in data_list:
        length = len(elem) // 2
        a, b = set(elem[:length]), set(elem[length:])
        total += get_v(list(a & b)[0])

    total = 0
    for gen in generator(data_list):
        inter = set(gen[0]) & set(gen[1]) & set(gen[2])
        total += get_v(list(inter)[0])
