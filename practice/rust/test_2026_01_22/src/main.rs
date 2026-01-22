fn bubble_sort(a: &mut [i32]) {
    let n = a.len();
    if n < 2 {
        return;
    }
    for i in 0..n {
        let mut swapped = false;
        for j in 0..(n - 1 - i) {
            if a[j] > a[j + 1] {
                a.swap(j, j + 1);
                swapped = true;
            }
        }
        if !swapped {
            break;
        }
    }
}

fn partition(a: &mut [i32]) -> usize {
    let len = a.len();
    let pivot = a[len - 1];
    let mut i = 0usize;
    for j in 0..(len - 1) {
        if a[j] < pivot {
            a.swap(i, j);
            i += 1;
        }
    }
    a.swap(i, len - 1);
    i
}

fn quick_sort(a: &mut [i32]) {
    let n = a.len();
    if n <= 1 {
        return;
    }
    let p = partition(a);
    quick_sort(&mut a[..p]);
    quick_sort(&mut a[p + 1..]);
}

fn main() {
    let mut v1 = vec![5, 2, 9, 1, 5, 6];
    let mut v2 = v1.clone();

    println!("原始: {:?}", v1);

    bubble_sort(&mut v1);
    println!("冒泡排序结果: {:?}", v1);

    quick_sort(&mut v2);
    println!("快速排序结果: {:?}", v2);
}
