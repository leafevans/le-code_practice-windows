class Program
{
    private const int INSERTION_SORT_CUTOFF = 10;

    static void QuickSort(int[] arr, int left, int right)
    {
        while (left < right)
        {
            if (right - left < INSERTION_SORT_CUTOFF)
            {
                InsertionSort(arr, left, right);
                return;
            }

            int pivotIndex = GetMedianOfThreePivotIndex(arr, left, right);
            int pivot = arr[pivotIndex];

            int i = left,
                j = right;
            while (i <= j)
            {
                while (arr[i] < pivot)
                    i++;
                while (arr[j] > pivot)
                    j--;
                if (i <= j)
                {
                    (arr[i], arr[j]) = (arr[j], arr[i]);
                    i++;
                    j--;
                }
            }

            if (j - left < right - i)
            {
                QuickSort(arr, left, j);
                left = i;
            }
            else
            {
                QuickSort(arr, i, right);
                right = j;
            }
        }
    }

    private static void InsertionSort(int[] arr, int left, int right)
    {
        for (int i = left + 1; i <= right; i++)
        {
            int current = arr[i];
            int j = i - 1;
            while (j >= left && arr[j] > current)
            {
                arr[j + 1] = arr[j];
                j--;
            }
            arr[j + 1] = current;
        }
    }

    private static int GetMedianOfThreePivotIndex(int[] arr, int left, int right)
    {
        int mid = left + (right - left) / 2;
        if (arr[left] > arr[mid])
            (arr[left], arr[mid]) = (arr[mid], arr[left]);
        if (arr[left] > arr[right])
            (arr[left], arr[right]) = (arr[right], arr[left]);
        if (arr[mid] > arr[right])
            (arr[mid], arr[right]) = (arr[right], arr[mid]);

        (arr[mid], arr[right - 1]) = (arr[right - 1], arr[mid]);
        return right - 1;
    }

    static void Main()
    {
        int[] arr = [5, 2, 9, 1, 5, 6, 12, 3, 8, 4, 10, 7];
        Console.WriteLine("排序前: " + string.Join(", ", arr));
        QuickSort(arr, 0, arr.Length - 1);
        Console.WriteLine("排序后: " + string.Join(", ", arr));
    }
}
