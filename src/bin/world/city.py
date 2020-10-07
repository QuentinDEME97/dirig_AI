class City:

    def __init__(self):
        self.stocks = {"iron" : 0, "wheat" : 0, "wood" : 0}

    def pay(self):
        self.stocks["wheat"] -= 5