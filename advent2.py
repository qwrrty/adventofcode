#! /usr/bin/python

# --- Day 2: I Was Told There Would Be No Math ---
#
# The elves are running low on wrapping paper, and so they need to
# submit an order for more. They have a list of the dimensions (length
# l, width w, and height h) of each present, and only want to order
# exactly as much as they need.
#
# Fortunately, every present is a box (a perfect right rectangular
# prism), which makes calculating the required wrapping paper for each
# gift a little easier: find the surface area of the box, which is
# 2*l*w + 2*w*h + 2*h*l. The elves also need a little extra paper for
# each present: the area of the smallest side.
#
# For example:
#
#   - A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52
#     square feet of wrapping paper plus 6 square feet of slack, for
#     a total of 58 square feet.
#   - A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42
#     square feet of wrapping paper plus 1 square foot of slack, for a
#     total of 43 square feet.
#
# All numbers in the elves' list are in feet. How many total square
# feet of wrapping paper should they order?


def dimensions(pkg_string):
    l, w, h = pkg_string.split('x')
    if not l or not w or not h:
        raise ValueError
    return int(l), int(w), int(h)


def pkg_area(pkg_string):
    l, w, h = dimensions(pkg_string)
    side1 = l * w
    side2 = w * h
    side3 = l * h
    return 2*side1 + 2*side2 + 2*side3 + min(side1, side2, side3)


def ribbon_length(pkg_string):
    l, w, h = dimensions(pkg_string)
    sides = sorted([l, w, h])
    return 2*sides[0] + 2*sides[1] + l*w*h


def main():
    with open("input2.txt") as f:
        pkg_list = f.readlines()

    total_area = sum([pkg_area(p) for p in pkg_list])
    print "area: ", total_area

    total_ribbon = sum([ribbon_length(p) for p in pkg_list])
    print "ribbon: ", total_ribbon


if __name__ == "__main__":
    main()
