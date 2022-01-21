reinter [![License](http://img.shields.io/:license-gpl3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0.html)
===

re2inter is built for checking two regular expressions have intersection (conflict) or not.
It's based on https://gitlab.com/opennota/re2dfa project.

# Installation

    go get -u github.com/oulinbao/reinter

# Usage

    make
    ./reinter "a*" "a+"

# License

re2dfa is released under the GNU General Public License version 3.0.  As a special exception to the GPLv3, you may use the parts of re2dfa output copied from re2dfa source without restriction.  Use of re2dfa makes no requirements about the license of generated code.

https://gitlab.com/opennota/re2dfa