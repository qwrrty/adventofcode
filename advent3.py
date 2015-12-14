#! /usr/bin/env python

# --- Day 3: Perfectly Spherical Houses in a Vacuum ---
#
# Santa is delivering presents to an infinite two-dimensional grid of houses.
#
# He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls
# him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v),
# east (>), or west (<). After each move, he delivers another present to the house at his new location.
#
# However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off,
# and Santa ends up visiting some houses more than once. How many houses receive at least one present?
#
# For example:
#
#   - > delivers presents to 2 houses: one at the starting location, and one to the east.
#   - ^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
#   - ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.

# --- Part Two ---
#
# The next year, to speed up the process, Santa creates a robot version of himself, Robo-Santa, to deliver presents
# with him.
#
# Santa and Robo-Santa start at the same location (delivering two presents to the same starting house),
# then take turns moving based on instructions from the elf, who is eggnoggedly reading from the same script as the
# previous year.
#
# This year, how many houses receive at least one present?
#
# For example:
#
#   - ^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.
#   - ^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.
#   - ^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.


import collections


class Santa(object):
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def move(self, dir):
        if dir == '<':
            self.x -= 1
        elif dir == '>':
            self.x += 1
        elif dir == '^':
            self.y -= 1
        elif dir == 'v':
            self.y += 1
        else:
            raise ValueError

    def housename(self):
        return "{},{}".format(self.x, self.y)


def yearone():
    with open("input3.txt") as f:
        directions = f.read()

    santa = Santa(0,0)
    presents = collections.defaultdict(int)
    presents[santa.housename()] += 1
    for dir in directions:
        santa.move(dir)
        presents[santa.housename()] += 1

    print "year 1: houses with at least one present: ", len(presents)


def yeartwo():
    with open("input3.txt") as f:
        directions = f.read()

    santa = Santa(0,0)
    robo = Santa(0,0)
    presents = collections.defaultdict(int)
    presents[santa.housename()] += 2
    i = 0
    for dir in directions:
        if i % 2 == 0:
            santa.move(dir)
            presents[santa.housename()] += 1
        else:
            robo.move(dir)
            presents[robo.housename()] += 1
        i += 1

    print "year 2: houses with at least one present: ", len(presents)


if __name__ == "__main__":
    yearone()
    yeartwo()
