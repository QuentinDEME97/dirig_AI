class Ruler:

    def __init__(self, personality):
        self.personality = personality

    def __repr__(self):
        representation = "Personality :"
        representation += "\nwar : "  + '|' * (self.personality['warrior']//10) + " (" + str(self.personality['warrior']) + ")"
        representation += "\ntru : "  + '|' * (self.personality['trust']//10) + " (" + str(self.personality['trust']) + ")"
        representation += "\nkin : "  + '|' * (self.personality['kindness']//10) + " (" + str(self.personality['kindness']) + ")"
        representation += "\ngre : "  + '|' * (self.personality['greed']//10) + " (" + str(self.personality['greed']) + ")" + "\n"
        return representation