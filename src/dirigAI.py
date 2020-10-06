from bin.bot.ruler_builder import RulerBuilder

def main():
    print("\n")
    builder = RulerBuilder()
    rulers = builder.generateN(3)
    
    for ruler in rulers:
        print(str(ruler)+"\n")

main()