from typing import Dict, List, Tuple
from utils import parse_input


def parse_input_polymers(problem_input: List[str]) -> Tuple[List[str], Dict[str, str]]:
    polymer = problem_input[0]
    rule_dict: Dict[str, str] = {}

    for rule in problem_input[2:]:
        k, v = rule.split(" -> ")
        rule_dict[k] = v

    return list(polymer), rule_dict


def solve1(problem_input: List[str], steps: int = 10) -> int:
    polymer, rule_dict = parse_input_polymers(problem_input)

    pair_counts: Dict[str, int] = {}

    for i in range(len(polymer)-1):
        pair = polymer[i] + polymer[i+1]
        pair_counts[pair] = pair_counts.get(pair, 0) + 1

    for _ in range(steps):
        pair_tuples = [(k, v) for k, v in pair_counts.items() if v > 0]
        for pair, count in pair_tuples:
            new_letter = rule_dict[pair]
            new_pair1 = pair[0] + new_letter
            new_pair2 = new_letter + pair[1]

            pair_counts[pair] = pair_counts.get(pair, 0) - count
            pair_counts[new_pair1] = pair_counts.get(new_pair1, 0) + count
            pair_counts[new_pair2] = pair_counts.get(new_pair2, 0) + count

    counts: Dict[str, int] = {}

    for k, v in pair_counts.items():
        counts[k[0]] = counts.get(k[0], 0) + v
        counts[k[1]] = counts.get(k[1], 0) + v

    # Account for the beginning and end only being in one pair
    counts[polymer[0]] += 1
    counts[polymer[-1]] += 1
    return (max(counts.values()) - min(counts.values())) // 2


def solve2(problem_input: List[str]) -> int:
    return solve1(problem_input, 40)


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day14.txt")

    # problem_input_ints = [int(num) for num in problem_input]

    # problem_input_2d_ints = [[int(num) for num in line]
    #                         for line in problem_input]

    print(solve1(problem_input))
    print(solve2(problem_input))
