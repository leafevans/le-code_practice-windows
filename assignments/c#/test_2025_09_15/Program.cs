using System;

class Program
{
    static void Main(string[] args)
    {
        int[] arr = [12, 34, 54, 2, 3];
        Console.WriteLine("Before: ");
        foreach (int num in arr)
        {
            Console.Write(num + " ");
        }
        HeapSort(arr);
        Console.WriteLine("\nAfter: ");
        foreach (int num in arr)
        {
            Console.Write(num + " ");
        }
    }

    static void Heapify(int[] arr, int n, int i)
    {
        int largest = i,
            left = 2 * i + 1,
            right = 2 * i + 2;

        if (left < n && arr[largest] < arr[left])
            largest = left;
        if (right < n && arr[largest] < arr[right])
            largest = right;

        if (largest != i)
        {
            (arr[largest], arr[i]) = (arr[i], arr[largest]);
            Heapify(arr, n, largest);
        }
    }

    static void HeapSort(int[] arr)
    {
        int n = arr.Length;

        for (int i = n / 2 - 1; i >= 0; --i)
            Heapify(arr, n, i);

        for (int i = n - 1; i > 0; --i)
        {
            (arr[0], arr[i]) = (arr[i], arr[0]);
            Heapify(arr, i, 0);
        }
    }

    static int[] MergeSort(int[] arr)
    {
        if (arr.Length <= 1)
            return arr;

        int mid = arr.Length / 2;
        int[] left = new int[mid];
        int[] right = new int[arr.Length - mid];

        Array.Copy(arr, 0, left, 0, mid);
        Array.Copy(arr, mid, right, 0, arr.Length - mid);

        left = MergeSort(left);
        right = MergeSort(right);

        return Merge(left, right);
    }

    static int[] Merge(int[] left, int[] right)
    {
        int[] result = new int[left.Length + right.Length];
        int i = 0,
            j = 0,
            k = 0;

        while (i < left.Length && j < right.Length)
        {
            if (left[i] < right[j])
                result[k++] = left[i++];
            else
                result[k++] = right[j++];
        }

        while (i < left.Length)
            result[k++] = left[i++];

        while (j < right.Length)
            result[k++] = right[j++];

        return result;
    }

    static void ShellSort(int[] arr)
    {
        int n = arr.Length;
        for (int gap = n / 2; gap > 0; gap /= 2)
        {
            for (int i = gap; i < n; ++i)
            {
                int key = arr[i];
                int j = i;
                while (j >= gap && arr[j - gap] > key)
                {
                    arr[j] = arr[j - gap];
                    j -= gap;
                }
                arr[j] = key;
            }
        }
    }

    static void InsertionSort(int[] arr)
    {
        int n = arr.Length;
        for (int i = 1; i < n; ++i)
        {
            int key = arr[i];
            int j = i - 1;
            while (j >= 0 && arr[j] > key)
            {
                arr[j + 1] = arr[j];
                --j;
            }
            arr[j + 1] = key;
        }
    }

    static void QuickSort(int[] arr, int left, int right)
    {
        if (left >= right)
            return;

        int pivot = arr[left + (right - left) / 2];
        int i = left;
        int j = right;
        while (i <= j)
        {
            while (arr[i] < pivot)
                ++i;
            while (arr[j] > pivot)
                --j;

            if (i <= j)
            {
                (arr[i], arr[j]) = (arr[j], arr[i]);
                ++i;
                --j;
            }
        }
        QuickSort(arr, left, i);
        QuickSort(arr, j, right);
    }

    static void BubbleSort(int[] arr)
    {
        int n = arr.Length;
        for (int i = 0; i < n - 1; ++i)
        for (int j = 0; j < n - 1 - i; ++j)
            if (arr[j + 1] < arr[j])
                (arr[j], arr[j + 1]) = (arr[j + 1], arr[j]);
    }

    static void SelectionSort(int[] arr)
    {
        int n = arr.Length;
        for (int i = 0; i < n; ++i)
        {
            int minIndex = i;
            for (int j = i + 1; j < n; ++j)
            {
                if (arr[j] < arr[minIndex])
                {
                    minIndex = j;
                }
            }
            if (minIndex != i)
            {
                (arr[minIndex], arr[i]) = (arr[i], arr[minIndex]);
            }
        }
    }
}
