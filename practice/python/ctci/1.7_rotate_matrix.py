def rotate(matrix):
    if len(matrix) == 0 or len(matrix) != len(matrix[0]):
        return False

    n = len(matrix)

    for layer in range(int(n/2)):
        first = layer
        last = n - 1 - layer

        for i in range(first, last):
            offset = i - first
            top = matrix[first][i]  # Save top

            # left -> top
            matrix[first][i] = matrix[last - offset][first]

            # bottom -> left
            matrix[last - offset][first] = matrix[last][last - offset]

            # right -> bottom
            matrix[last][last - offset] = matrix[first + offset][last]

            # top -> right
            matrix[first + offset][last] = top

    return True


if __name__ == "__main__":
    matrix = [
        [0, 0, 0, 0],
        [1, 1, 1, 1],
        [2, 2, 2, 2],
        [3, 3, 3, 3],
    ]

    rotate(matrix)

    for row in matrix:
        print(row)
