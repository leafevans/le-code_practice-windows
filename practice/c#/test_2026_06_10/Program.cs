namespace test_2026_06_10
{
    public static class QuickSort
    {
        public static void Sort<T>(T[] array, int left, int right)
           where T : IComparable<T>
        {
            if (left >= right) return;
            int pivotIndex = Partition(array, left, right);
            Sort(array, left, pivotIndex - 1);
            Sort(array, pivotIndex + 1, right);
        }

        private static int Partition<T>(T[] array, int left, int right)
            where T : IComparable<T>
        {
            int mid = left + (right - left) / 2;
            if (array[left].CompareTo(array[mid]) > 0) Swap(array, left, mid);
            if (array[left].CompareTo(array[right]) > 0) Swap(array, left, right);
            if (array[mid].CompareTo(array[right]) > 0) Swap(array, mid, right);
            Swap(array, mid, right);
            T pivot = array[right];
            int i = left - 1;
            for (int j = left; j < right; j++)
            {
                if (array[j].CompareTo(pivot) <= 0)
                {
                    ++i;
                    Swap(array, i, j);
                }
            }
            Swap(array, i + 1, right);
            return i + 1;
        }

        private static void Swap<T>(T[] array, int a, int b)
        {
            (array[a], array[b]) = (array[b], array[a]);
        }

        public static void Main()
        {
            int[] arr = [9, 3, 7, 1, 5, 8, 2, 6, 4];
            Console.WriteLine("排序前: " + string.Join(", ", arr));
            Sort(arr, 0, arr.Length - 1);
            Console.WriteLine("排序后: " + string.Join(", ", arr));
        }
    }
}
