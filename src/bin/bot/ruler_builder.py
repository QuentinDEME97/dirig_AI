from .ruler import Ruler
import random

class RulerBuilder:

    def __init__(self):
        pass

    def generateOne(self):
        ruler_personality = {}
        ruler_personality['warrior'] = int(round(random.random(), 2) * 100)
        ruler_personality['trust'] = int(round(random.random(), 2) * 100)
        ruler_personality['kindness'] = int(round(random.random(), 2) * 100)
        ruler_personality['greed'] = int(round(random.random(), 2) * 100)
        return Ruler(ruler_personality)
    
    def generateN(self, n):
        rulers = []
        for i in range(n):
            rulers.append(self.generateOne())
        return rulers