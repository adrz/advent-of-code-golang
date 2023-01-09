from utils import download_input

if __name__ == "__main__":
    data = download_input(2)
    my_dict = {
        "Y": {
            "A": 2 + 6,
            "B": 2 + 3,
            "C": 2,
        },
        "X": {
            "A": 1 + 3,
            "B": 1,
            "C": 1 + 6,
        },
        "Z": {
            "A": 3,
            "B": 3 + 6,
            "C": 3 + 3,
        },
    }

    rounds = data.split("\n")
    total_score = 0
    for r in rounds[:-1]:
        k1, k2 = r.split(" ")
        total_score += my_dict[k2][k1]

    my_dict = {
        "Y": {
            "A": 1 + 3,
            "B": 2 + 3,
            "C": 3 + 3,
        },
        "X": {
            "A": 0 + 3,
            "B": 0 + 1,
            "C": 0 + 2,
        },
        "Z": {
            "A": 6 + 2,
            "B": 6 + 3,
            "C": 6 + 1,
        },
    }

    rounds = data.split("\n")
    total_score = 0
    for r in rounds[:-1]:
        k1, k2 = r.split(" ")
        total_score += my_dict[k2][k1]
