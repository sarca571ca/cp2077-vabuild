## Virtual Atelier Store Builder

### Store ID (Character Only, no space or special characters)
Example: CoolAssStore
### Store Name (What is seen in game)
Example: Cool Ass Store
### Atlas Resource
I'm assuming this is a special graphic file .inkatlas. Just put the path
to the file here.
### Texture Type (Assuming this is for the Class of Texture?)
Example: CHARACTER

### Header for the .reds file
@addMethod(gameuiInGameMenuGameController)
protected cb func RegisterTheDreamShopStore(event: ref<VirtualShopRegistration>) -> Bool {
  event.AddStore(
    n"CoolAssStore",
    "Cool Ass Store",
    [
    "Items.Preset_Ajax_Default",
    "Items.Preset_Ajax_Military",
    "Items.Preset_Ajax_Moron",
    "Items.Preset_Ajax_Neon",
    "Items.Preset_Ajax_Pimp",
    "Items.Preset_Ajax_Training",
    "Items.Preset_Base_Copperhead"
    ],
    [], // Money Value - Leave blank to let it scale with player
    r"base/gameplay/gui/world/adverts/naranjita/naranjita_atlas.inkatlus",
    n"CHARACTER",
    [] // Quality (Common, Uncommon, Rare, Iconic, Legendary, Random) - Leave it blank to be random
  );
}

### Installation Locations
{GAMEDIR}\r6\scripts\
