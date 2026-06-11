namespace test_2026_06_10
{
    public static class QuickSort
    {
        // 快速排序函数
        public static void Sort<T>(T[] array, int left, int right)
           where T : IComparable<T>
        {
            // 如果左边界大于等于右边界，直接返回
            if (left >= right) return;

            int pivotIndex = Partition(array, left, right);
            Sort(array, left, pivotIndex - 1);
            Sort(array, pivotIndex + 1, right);
        }

        private static int Partition<T>(T[] array, int left, int right)
            where T : IComparable<T>
        {
            // 选择中间元素作为基准值
            int mid = left + (right - left) / 2;
            if (array[left].CompareTo(array[mid]) > 0) Swap(array, left, mid);
            if (array[left].CompareTo(array[right]) > 0) Swap(array, left, right);
            if (array[mid].CompareTo(array[right]) > 0) Swap(array, mid, right);
            // 将基准值移动到数组末尾
            Swap(array, mid, right);
            T pivot = array[right];
            int i = left - 1;
            // 遍历数组，将小于等于基准值的元素移动到基准值的左边
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

        // 交换数组中两个元素的位置
        private static void Swap<T>(T[] array, int a, int b)
        {
            (array[a], array[b]) = (array[b], array[a]);
        }

        // 快速排序主函数
        public static void Main()
        {
            int[] arr = [9, 3, 7, 1, 5, 8, 2, 6, 4];
        }
    }
}
