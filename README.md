# Plastic Slate

Original: _[A simple syntax and UI theme for Visual Studio Code.](http://wstone.io/plastic/)_

---

Modification of will-stone's Plastic for the iTerm2 Slate color palette. Semi-bright pastels replace bold.

## Changes over Plastic

**Plastic Original**  
<img src="https://raw.githubusercontent.com/kayteh/plastic-slate/master/docs/screenshot-orig.png" width="500" alt="Screenshot">

**Plastic Slate**  
<img src="https://raw.githubusercontent.com/kayteh/plastic-slate/master/docs/screenshot-slate.png" width="500" alt="Screenshot">

## Colours

Colours are taken from the [iTerm2 Slate](https://github.com/mbadolato/iTerm2-Color-Schemes#slate).

Most are left alone from the original Plastic scheme, but may get adjusted in the future.

This theme also includes colour changes to the Integrated Terminal to keep it in line with the colour source.

Colour names provided by [Chirag Mehta's](http://chir.ag/about) fantastic [Name That Colour](http://chir.ag/projects/name-that-color/) website.


### Syntax

* ![#c4c9c0](https://placehold.it/15/c4c9c0/000000?text=+) Kangaroo **#c4c9c0** - _text & variables_
* ![#c5a7d9](https://placehold.it/15/c5a7d9/000000?text=+) Light Wisteria **#c5a7d9** - _storage_
* ![#15ab9c](https://placehold.it/15/15ab9c/000000?text=+) Mountain Meadow **#15ab9c** - _constant_
* ![#5c6370](https://placehold.it/15/5c6370/000000?text=+) Shuttle Gray **#5c6370** - _comments_
* ![#648890](https://placehold.it/15/648890/000000?text=+) Juniper **#648890** - _deprioritized punctuation_
* ![#d0ccca](https://placehold.it/15/d0ccca/000000?text=+) Cloud **#d0ccca** - _prioritized punctuation_
* ![#81d778](https://placehold.it/15/81d778/000000?text=+) Olivine **#81d778** - _strings_
* ![#7ab0d2](https://placehold.it/15/7ab0d2/000000?text=+) Half Baked **#7ab0d2** - _functions_
* ![#a481d3](https://placehold.it/15/a481d3/000000?text=+) Lilac Blush **#a481d3** - _tags_
* ![#e2a8bf](https://placehold.it/15/e2a8bf/000000?text=+) Light Orchid **#e2a8bf** - _keywords_

### UI

* ![#1B659D](https://placehold.it/15/1B659D/000000?text=+) Matisse **#1B659D** - _Info_
* ![#A86515](https://placehold.it/15/d4ac0d/000000?text=+) Galliano **#d4ac0d** - _Warnings_
* ![#C62D42](https://placehold.it/15/C62D42/000000?text=+) Brick Red **#C62D42** - _Errors_
* ![#3E4451](https://placehold.it/15/3E4451/000000?text=+) Bright Gray **#3E4451** - _Line numbers, invisibles, ruler_
* ![#21252b](https://placehold.it/15/21252b/000000?text=+) Shark **#21252b** - _Fore-background (activity bar, status bar, editor)_
* ![#181A1F](https://placehold.it/15/181A1F/000000?text=+) Woodsmoke **#181A1F** - _Back-background (side bar)_
* ![#080A0F](https://placehold.it/15/080A0F/000000?text=+) Bunker **#080A0F** - _Back-back-background (input background, contrast borders)_


## Contribute

Themes can always be improved and tweaked so please feel free to log issues and send pull requests. However, please keep in mind that this theme strives to remain simple and language agnostic; for example, please do not ask for `blah.blah.php` or `blah.blah.c` scopes to be added.

**Fork specific note:** Some changes are better left to upstream Plastic, and merged into this repo. Color modifications, documentation, and tooling are the only pull requests recommended here.

Clone the repository to `.vscode/extenstions/`(remember to remove any other installed versions of Plastic Slate), run `yarn` to install dependencies, and `yarn docs` to serve the documentation page.

Use Microsoft's [theme color reference](https://code.visualstudio.com/docs/getstarted/theme-color-reference) for UI theming help.

You'll notice that the syntax is grouped by colour instead of scope. I find this much nicer to maintain.

Happy hacking :)

# License

MIT
