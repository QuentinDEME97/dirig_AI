from bin.game_factory import GameFactory

def main():
    print("\n")
    builder = GameFactory()
    rulers = builder.generateNRuler(3)
    world = builder.generateWorld(5,5)
    print(world.grid)
    
    for ruler in rulers:
        print(str(ruler)+"\n")

main()