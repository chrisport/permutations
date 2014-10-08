package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPermutations(t *testing.T) {
	Convey("Subject: Test generation of permutations of string, tail recursive implementation", t, func() {

			Convey("Test Permutations with one char, should contain only character itself", func() {
					oneCharString := "a"
					result := permutations(oneCharString)
					So(len(result), ShouldEqual, 1)
					So(result[0], ShouldEqual, "a")
				})
			Convey("Test Permutations with two chars, should contain original and swapped version", func() {
					twoCharString := "ab"
					result := permutations(twoCharString)
					So(len(result), ShouldEqual, 2)
					So(result, ShouldContain, "ab")
					So(result, ShouldContain, "ba")
				})
			Convey("Test Permutations with three chars, should contain all 6 combinations", func() {
					threeCharString := "abc"
					result := permutations(threeCharString)
					So(len(result), ShouldEqual, 6)
					So(result, ShouldContain, "abc")
					So(result, ShouldContain, "acb")
					So(result, ShouldContain, "cab")
					So(result, ShouldContain, "cba")
					So(result, ShouldContain, "bac")
					So(result, ShouldContain, "bca")
				})
			Convey("Test Permutations with four, count of permutations should be equals factorial of length of inputstring", func() {
					fourCharString := "abcd"
					result := permutations(fourCharString)
					So(len(result), ShouldEqual, 24)

					fiveCharString := "abcde"
					result = permutations(fiveCharString)
					So(len(result), ShouldEqual, 120)

					sixCharString := "abcdef"
					result = permutations(sixCharString)
					So(len(result), ShouldEqual, 720)
				})
		})
}

func TestPermutationsTailRecursive(t *testing.T) {
	Convey("Subject: Test generation of permutations of string, tail recursive implementation", t, func() {

			Convey("Test Permutations with one char, should contain only character itself", func() {
					oneCharString := "a"
					result := permutationsTailRecursive(oneCharString)
					So(len(result), ShouldEqual, 1)
					So(result[0], ShouldEqual, "a")
				})
			Convey("Test Permutations with two chars, should contain original and swapped version", func() {
					twoCharString := "ab"
					result := permutationsTailRecursive(twoCharString)
					So(len(result), ShouldEqual, 2)
					So(result, ShouldContain, "ab")
					So(result, ShouldContain, "ba")
				})
			Convey("Test Permutations with three chars, should contain all 6 combinations", func() {
					threeCharString := "abc"
					result := permutationsTailRecursive(threeCharString)
					So(len(result), ShouldEqual, 6)
					So(result, ShouldContain, "abc")
					So(result, ShouldContain, "acb")
					So(result, ShouldContain, "cab")
					So(result, ShouldContain, "cba")
					So(result, ShouldContain, "bac")
					So(result, ShouldContain, "bca")
				})
			Convey("Test Permutations with four, count of permutations should be equals factorial of length of inputstring", func() {
					fourCharString := "abcd"
					result := permutationsTailRecursive(fourCharString)
					So(len(result), ShouldEqual, 24)

					fiveCharString := "abcde"
					result = permutations(fiveCharString)
					So(len(result), ShouldEqual, 120)

					sixCharString := "abcdef"
					result = permutations(sixCharString)
					So(len(result), ShouldEqual, 720)
				})
		})
}
