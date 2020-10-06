from .ruler import Ruler
import random
import codecs

class RulerBuilder:

    def __init__(self):
        with codecs.open("bin/resources/names.txt", encoding='utf-8') as f:
            self.names_list = f.read().split(",")
            f.close()
        with codecs.open("bin/resources/adjectives.txt", encoding='utf-8') as f:
            self.adjectives_list = f.read().split(",")
            f.close()

    def generateOne(self):
        ruler_personality = {}
        ruler_personality['warrior'] = int(round(random.random(), 2) * 100)
        ruler_personality['trust'] = int(round(random.random(), 2) * 100)
        ruler_personality['kindness'] = int(round(random.random(), 2) * 100)
        ruler_personality['greed'] = int(round(random.random(), 2) * 100)
        
        name = self.generateName()

        return Ruler(name, ruler_personality)
    
    def generateN(self, n):
        rulers = []
        for i in range(n):
            rulers.append(self.generateOne())
        return rulers

    def generateName(self):
        name = random.choice(self.names_list)
        adjective = random.choice(self.adjectives_list)
        self.names_list.remove(name)
        self.adjectives_list.remove(adjective)
        return name + adjective
        
    
