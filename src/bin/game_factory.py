from .bot.ruler import Ruler
from .world import World, City
import random
import codecs

class GameFactory:

    def __init__(self):
        with codecs.open("bin/resources/names.txt", encoding='utf-8') as f:
            self.names_list = f.read().split(",")
            f.close()
        with codecs.open("bin/resources/adjectives.txt", encoding='utf-8') as f:
            self.adjectives_list = f.read().split(",")
            f.close()

    def generateOneRuler(self):
        ruler_personality = {}
        ruler_personality['warrior'] = round(random.random(), 2)
        ruler_personality['trust'] = round(random.random(), 2)
        ruler_personality['kindness'] = round(random.random(), 2)
        ruler_personality['greed'] = round(random.random(), 2)
        
        name = self.generateName()

        ruler = Ruler(name, ruler_personality)
        ruler.addCity(City((5,5)))

        return ruler
    
    def generateNRuler(self, n):
        rulers = []
        for i in range(n):
            rulers.append(self.generateOneRuler())
        return rulers

    def generateName(self):
        name = random.choice(self.names_list)
        adjective = random.choice(self.adjectives_list)
        self.adjectives_list.remove(adjective)
        return name + adjective

    def generateWorld(self, sizeX, sizeY):
        grid = []
        for x in range(sizeX):
            grid.append([])
            for y in range(sizeY):
                grid[x].append('_')
        return World(grid, {})
        
        
    
