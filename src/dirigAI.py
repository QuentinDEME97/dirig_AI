from bin.game_factory import GameFactory

def main():
    print("\n")
    builder = GameFactory()
    ruler = builder.generateOneRuler()
    world = builder.generateWorld(5,5)
    print(world.grid)
    
    print(ruler.personality)

    print(ruler)
    print("Stock low" if ruler.city.isStockLow() else "Tout va bien " + str(ruler.city.stocks["wheat"]) + " blé")
    print("Should buy ? " + str(ruler.shouldBuyWheat()))
    ruler.pay()
    print("Stock low " + str(ruler.city.stocks["wheat"]) + " blé" if ruler.city.isStockLow() else "Tout va bien " + str(ruler.city.stocks["wheat"]) + "blé")
    print("Should buy ? " + str(ruler.shouldBuyWheat()))
    
main()