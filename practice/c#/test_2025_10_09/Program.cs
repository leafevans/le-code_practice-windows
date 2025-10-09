// using System;

// class Program
// {
//     static void PrintJaggedArrayAddresses()
//     {
//         int[][] jaggedArray =
//         [
//             [0, 1, 2],
//             [3, 4, 5],
//             [6, 7, 8],
//             [9, 9, 9],
//         ];

//         for (int row = 0; row < jaggedArray.Length; row++)
//         {
//             System.Console.WriteLine(
//                 $"第 {row} 行数组对象地址：{jaggedArray[row].GetHashCode():X}"
//             );
//         }
//     }

//     static void Main(string[] args)
//     {
//         PrintJaggedArrayAddresses();
//     }
// }
using System;

class Program
{
    static void TestRectArr()
    {
        int[,] arr = new int[,]
        {
            { 0, 1, 2 },
            { 3, 4, 5 },
            { 6, 7, 8 },
            { 9, 9, 9 },
        };

        for (int i = 0; i < arr.GetLength(0); i++)
        {
            for (int j = 0; j < arr.GetLength(1); j++)
            {
                System.Console.Write(arr[i, j] + " ");
            }
            System.Console.WriteLine();
        }
    }

    static void Main(string[] args)
    {
        TestRectArr();
    }
}
