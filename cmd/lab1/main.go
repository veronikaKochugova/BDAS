package main

import "fmt"

const (
	source = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	target = "Q5A8ZWS0XEDC6RFVT9GBY4HNU3J2MI1KO7LP"
)

func Obfuscate(s string) (result string) {
	// public static String obfuscate(String s) {
	// 	char[] result= new char[10];
	// 	for (int i=0;i<s.length();i++) {
	// 	char c=s.charAt(i);
	// 	int index=source.indexOf(c);
	// 	result[i]=target.charAt(index);
	// 	}
	// 	return new String(result);
	// 	}
	return source
}

func Unobfuscate(s string) (result string) {
	// 	public static String unobfuscate(String s) {
	// 	char[] result= new char[10];
	// 	for (int i=0;i<s.length();i++) {
	// 	char c=s.charAt(i);
	// 	int index=target.indexOf(c);
	// 	result[i]=source.charAt(index);
	// 	}
	// 	return new String(result);
	// 	}
	return target
}

func main() {
	fmt.Println("Hello")
}
