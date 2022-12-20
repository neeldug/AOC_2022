use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::collections::BinaryHeap;
use std::num::ParseIntError;


fn main() -> Result<(), ParseIntError> {
    let elf_calories = build_elf_heap()?;
    part_1(&elf_calories);
    part_2(elf_calories);

    Ok(())
}

fn part_1(elf_calories: &BinaryHeap<usize>) {
    if let Some(c) = elf_calories.peek() {
        println!("{}", c);
    } else {
        eprintln!("No elfs in binary heap");
    }
}

fn part_2(mut elf_calories: BinaryHeap<usize>) {
    let mut acc = 0;
    for _ in 0..3 {
        if let Some(v) = elf_calories.pop() {
            acc += v;
        }
    }
    println!("{}", acc);
}

fn build_elf_heap() -> Result<BinaryHeap<usize>, ParseIntError> {
    let mut elf_calories = BinaryHeap::new();
    if let Ok(lines) = read_lines("input.txt") {
        let mut acc = 0usize;
        for line in lines {
            if let Ok(l) = line {
                if l.is_empty() {
                    elf_calories.push(acc);
                    acc = 0;
                } else {
                    let val = l.parse::<usize>()?;
                    acc += val;
                }
            }
        }
    }
    Ok(elf_calories)
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}