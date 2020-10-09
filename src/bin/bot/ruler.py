class Ruler:

    def __init__(self, name, personality):
        self.name = name
        self.personality = personality
    
    def addCity(self, city):
        self.city = city

    def pay(self):
        self.city.pay()

    def shouldBuyWheat(self):
        stockState = 1 if self.city.isStockLow() else 0
        return stockState + (self.personality["greed"]/2) + (self.personality["warrior"]/3)

    def __repr__(self):
        representation = self.name + " :"
        representation += "\nwar : "  + '|' * (int(self.personality['warrior']*100)) + " (" + str(self.personality['warrior']) + ")"
        representation += "\ntru : "  + '|' * (int(self.personality['trust']*100)) + " (" + str(self.personality['trust']) + ")"
        representation += "\nkin : "  + '|' * (int(self.personality['kindness']*100)) + " (" + str(self.personality['kindness']) + ")"
        representation += "\ngre : "  + '|' * (int(self.personality['greed']*100)) + " (" + str(self.personality['greed']) + ")" + "\n"
        representation += "Position de la cité : " + str(self.city.position)
        return representation