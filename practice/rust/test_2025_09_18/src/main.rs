// fn selection_sort(arr: &mut [i32]) {
//     let n = arr.len();
//     for i in 0..n {
//         let mut min_index = i;
//         for j in (i + 1)..n {
//             if arr[j] < arr[min_index] {
//                 min_index = j;
//             }
//             arr.swap(i, min_index)
//         }
//     }
// }

// fn bubble_sort(arr: &mut [i32]) {
//     let n = arr.len();
//     for i in 0..(n - 1) {
//         for j in 0..(n - 1 - i) {
//             if arr[j + 1] < arr[j] {
//                 arr.swap(j + 1, j)
//             }
//         }
//     }
// }

// fn insertion_sort(arr: &mut [i32]) {
//     let n = arr.len();
//     for i in 1..n {
//         let mut j: isize = (i - 1) as isize;
//         let key = arr[i];
//         while j >= 0 && arr[j as usize] > key {
//             arr[(j + 1) as usize] = arr[j as usize];
//             j -= 1;
//         }
//         arr[(j + 1) as usize] = key;
//     }
// }

fn quick_sort(arr: &mut [i32]) {
    if arr.len() <= 1 {
        return;
    }

    let pivot = arr[arr.len() / 2];
    let mut left: isize = 0;
    let mut right: isize = (arr.len() - 1) as isize;

    while left <= right {
        while left < arr.len() as isize && arr[left as usize] < pivot {
            left += 1;
        }

        while right >= 0 && arr[right as usize] > pivot {
            right -= 1;
        }

        if left <= right {
            arr.swap(left as usize, right as usize);
            left += 1;
            right -= 1;
        }
    }

    if right >= 0 {
        quick_sort(&mut arr[0..=(right as usize)]);
    }
    if left < arr.len() as isize {
        quick_sort(&mut arr[(left as usize)..]);
    }
}

fn merge(arr: &mut [i32], mid: usize) {
    let left = arr[0..mid].to_vec();
    let right = arr[mid..].to_vec();
    let mut i = 0;
    let mut j = 0;
    let mut k = 0;
    
    while i < left.len() && j < right.len() {
        if left[i] <= right[j] {
            arr[k] = left[i];
            i += 1;
        } else {
            arr[k] = right[j];
            j += 1;
        }
        k += 1;
    }

    while i < left.len() {
        arr[k] = left[i];
        i += 1;
        k += 1;
    }

    while j < right.len() {
        arr[k]  = right[j];
        j += 1;
        k += 1;
    }
 }

 fn merge_sort(arr :&mut[i32]) {
    if arr.len() <= 1 {
        return;
    }

    let mid = arr.len() / 2;

    merge_sort(&mut arr[0..mid]);
    merge_sort(&mut arr[mid..]);

    merge(arr, mid);
 }

fn main() {
    let mut arr = vec![64, 25, 12, 22, 11];
    println!("原始数组: {:?}", arr);
    merge_sort(&mut arr);
    println!("排序后数组：{:?}", arr);
}
