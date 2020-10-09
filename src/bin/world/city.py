class City:

    def __init__(self, position):
        self.stocks = {"iron" : 0, "wheat" : 11, "wood" : 0}
        self.position = position

    def pay(self):
        self.stocks["wheat"] -= 5

    def isStockLow(self):
        return self.stocks["wheat"] - 5 <= 5