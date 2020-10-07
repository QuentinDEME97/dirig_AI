from .bot.ruler import Ruler
from .world.world import World
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
        ruler_personality['warrior'] = int(round(random.random(), 2) * 100)
        ruler_personality['trust'] = int(round(random.random(), 2) * 100)
        ruler_personality['kindness'] = int(round(random.random(), 2) * 100)
        ruler_personality['greed'] = int(round(random.random(), 2) * 100)
        
        name = self.generateName()

        return Ruler(name, ruler_personality)
    
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
        
        
    
