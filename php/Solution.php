<?php

class Solution {

    /**
     * @param String $s
     * @param String $t
     * @return Boolean
     */
    function isAnagram($s, $t) {

        if ($s == $t) {
            return true;
        }

        if (strlen($s) != strlen($t)) {
            return false;
        }

        $firstString = str_split($s);
        $secondString = str_split($t);

        sort($firstString, \SORT_STRING);
        sort($secondString, \SORT_STRING);

        return $firstString == $secondString;
    }
}

$solution = new Solution();
echo (int) $solution->isAnagram('anagram', 'nagaram') . "\n";
echo (int) $solution->isAnagram('rat', 'car') . "\n";
