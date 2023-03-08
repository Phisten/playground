for x6 in range(1, 56):
    x11 = 99 - x6 - 1
    x1 = 99 - x6 - x11
    x2 = 99 - x1 - x6
    x3 = 99 - x2 - x1
    x4 = 99 - x3 - x2
    x5 = 99 - x4 - x3
    x7 = 99 - x6 - x5
    x8 = 99 - x7 - x6
    x9 = 99 - x8 - x7
    x10 = 99 - x9 - x8
    if len(set([x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11])) == 11 and \
            all(1 <= x <= 55 for x in [x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11]):
        print(x6)
        break
