board_size = 7
board = []
for i in range(7):
    board.append([])
    for _ in range(7):
        board[i].append('.')


def print_board(board):
    for r in board:
        print(''.join(r))
    print()


def get_tile(player):
    if player == 0:
        return 'W'
    return 'B'


def get_neighbors(board, x, y):
    neighbors = []
    if x-1 >= 0:
        neighbors.append((x-1,y))
    if y-1 >= 0:
        neighbors.append((x,y-1))
    if y+1 < len(board[x]):
        neighbors.append((x,y+1))
    if x+1 < len(board):
        neighbors.append((x+1, y))
    return neighbors


# capture if necessay will accept a board state and the most recently played
# move (x, y). will then validate whether degrees of freedom of contiguous
# opponent pieces have been reduced to 0. If so, remove pieces (and
# eventually update score)
def capture_if_necessary(board, x, y):
    capturing_player = board[x][y]
    neighbors = get_neighbors(board, x, y)
    for n in neighbors:
        bfs(board, n, capturing_player)


# BFS from starting position x, y -- capturing visited nodes if necessary
def bfs(board, pos, capturing_player):
    visited = set()
    stack = [(pos[0],pos[1])]
    while len(stack) != 0:
        curr_x, curr_y = stack[0]
        neighbors = get_neighbors(board, curr_x, curr_y)
        for n in neighbors:
            if n in visited:
                continue
            square = board[n[0]][n[1]]
            if square == '.':
                return
            if square == capturing_player:
                continue
            if square == 'W':
                stack.append(n)
        visited.add((curr_x, curr_y))
        stack = stack[1:]
    for v in visited:
        board[v[0]][v[1]] = '.'


def play_move(board, x, y, player):
    if board[x][y] != '.':
        return False
    board[x][y] = get_tile(player)
    capture_if_necessary(board, x, y)
    return True


player = 0
while True:
    print_board(board)
    move = input("player {} where would you like to place a tile".format(player))
    if move == "quit":
        break
    while True:
        # move format: i,j
        coords = move.split(',')
        x, y = int(coords[0]), int(coords[1])
        played_move = play_move(board, x, y, player)
        if played_move:
            break
        move = input("invalid move please try again")
    player = (player + 1) % 2


