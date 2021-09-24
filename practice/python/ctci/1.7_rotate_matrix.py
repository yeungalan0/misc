def rotate(matrix):
    if len(matrix) == 0 or len(matrix) != len(matrix[0]):
        return False

    n = len(matrix)

    for layer in range(n//2):
        first = layer
        last = n - 1 - layer

        for i in range(first, last):
            offset = i - first # Which index is going to be rotated
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

# Second attempt...
def rotate2(matrix):
    start = 0
    end = len(matrix) - 1

    while start < end:
        for i in range(start, end):
            offset = end - i + start # Mistake this needs to shrink by i but grow by start to always hit the end of the matrix layer

            temp = matrix[offset][start] # get left
            temp_right = matrix[i][end]
            matrix[i][end] = matrix[start][i] # top -> right
            temp_bottom = matrix[end][offset]
            matrix[end][offset] = temp_right # right -> bottom
            matrix[offset][start] = temp_bottom # bottom -> left
            matrix[start][i] = temp # left -> top

        start += 1
        end -= 1

        print("NEW ITERATION:")
        for row in matrix:
            print(row)

    return matrix


if __name__ == "__main__":
    matrix = [
        [7, 0, 0, 8],
        [1, 1, 1, 1],
        [2, 2, 2, 2],
        [6, 3, 3, 9],
    ]

    # matrix = [
    #     [1, 1],
    #     [2, 2],
    # ]

    rotate2(matrix)

    print("Result:")
    for row in matrix:
        print(row)
