# ani-cli-dw
## Download wrapper for https://github.com/pystardust/ani-cli
ani-cli-dw enhances ani-cli functionality by enabling simultaneous downloading of multiple episodes, which isn't an option in the current version of ani-cli.

## How to use
To use ani-cli-dw, follow these steps:

1. Run `ani-cli`.
2. Search for the show you want to download and remember its name, e.g. `Lupin III`.
3. Select the entry and remember its ID, e.g. `39`.
4. Exit ani-cli with ESC or `Ctrl-C`.
5. Run `ani-cli-dw`.
6. Enter the show name, e.g. `Lupin III`.
7. Enter the entry ID, e.g. `39`.
8. Follow the rest of the prompts to initiate the download.

## Requirements
- Windows (for now) as it runs cmd.exe  
- [github.com/pystardust/ani-cli](https://github.com/pystardust/ani-cli) installed and running 

## How to build
simply run `go build`

## Issues you may encounter
- Some episodes may take significantly longer to download.
  - This could occur due to large file sizes or exceeding your connection capacity. Consider reducing the number of simultaneous downloads.
- Errors may occur during downloading.
  - These errors could be caused by exceeding connection capabilities or host restrictions. The application will notify you of failed episodes. Simply rerun the same range; it will skip already downloaded episodes and attempt to download the failed ones again.
  - Can be also caused if you try to download non-existing episodes
