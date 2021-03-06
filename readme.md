#GolangSortingVisualization

[![Coverage Status](https://coveralls.io/repos/SimonWaldherr/GolangSortingVisualization/badge.png)](https://coveralls.io/r/SimonWaldherr/GolangSortingVisualization) 
[![Build Status](https://travis-ci.org/SimonWaldherr/GolangSortingVisualization.svg?branch=master)](https://travis-ci.org/SimonWaldherr/GolangSortingVisualization) 
[![Flattr donate button](https://raw.github.com/balupton/flattr-buttons/master/badge-89x18.gif)](https://flattr.com/submit/auto?user_id=SimonWaldherr&url=http%3A%2F%2Fgithub.com%2FSimonWaldherr%2FGolangSortingVisualization "Donate monthly to this project using Flattr")


##Sort Algorithms

###BogoSort

![Bubble Sort Animation](http://simonwaldherr.github.io/GolangSortingVisualization/sort_bogo.gif)

###BubbleSort

![Bubble Sort Animation](http://simonwaldherr.github.io/GolangSortingVisualization/sort_bubble.gif)

###CombSort

![Comb Sort Animation](http://simonwaldherr.github.io/GolangSortingVisualization/sort_comb.gif)

###CountingSort

![Counting Sort Animation](http://simonwaldherr.github.io/GolangSortingVisualization/sort_counting.gif)

###GnomeSort

![Gnome Sort Animation](http://simonwaldherr.github.io/GolangSortingVisualization/sort_gnome.gif)

###InsertionSort

![Insertion Sort Animation](http://simonwaldherr.github.io/GolangSortingVisualization/sort_insertion.gif)

###OddEvenSort

![OddEven Sort Animation](http://simonwaldherr.github.io/GolangSortingVisualization/sort_oddEven.gif)

###SelectionSort

![Selection Sort Animation](http://simonwaldherr.github.io/GolangSortingVisualization/sort_selection.gif)

###SleepSort

![Sleep Sort Animation](http://simonwaldherr.github.io/GolangSortingVisualization/sort_sleep.gif)

###StoogeSort

![Stooge Sort Animation](http://simonwaldherr.github.io/GolangSortingVisualization/sort_stooge.gif)

##HowTo

```sh
./start.sh
```

```sh
$ go run gsv.go --help
Usage of gsv:
  -algo="bubble": Select sorting algorithm all/bogo/[bubble]/comb/counting/gnome/insertion/oddEven/selection/sleep
  -count=30: number of values
  -fps=10: frames per second
  -max=9: highest value
  -mode=1: visualization mode
  -vis="stdout": Select output: [stdout]/gif
```

##License

MIT