from bin.bot.ruler_builder import RulerBuilder

def main():
    builder = RulerBuilder()
    ruler = builder.generateOne()
    
    print(ruler)

main()