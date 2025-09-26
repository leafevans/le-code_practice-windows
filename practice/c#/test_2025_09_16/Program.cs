class Program
{
    static void Main(string[] args)
    {
        string text = "ABABCABCACBAB";
        string pattern = "ABCAC";

        Console.WriteLine(KmpSearch(text, pattern));
    }

    static int[] BuildNext(string pattern)
    {
        int[] next = new int[pattern.Length];
        int j = 0;

        for (int i = 1; i < pattern.Length; ++i)
        {
            while (j > 0 && pattern[i] != pattern[j])
                j = next[j - 1];

            if (pattern[i] == pattern[j])
                ++j;

            next[i] = j;
        }

        return next;
    }

    static int KmpSearch(string text, string pattern)
    {
        if (pattern.Length == 0)
            return 0;

        int[] next = BuildNext(pattern);
        int j = 0;

        for (int i = 0; i < text.Length; ++i)
        {
            while (j > 0 && text[i] != pattern[j])
                j = next[j - 1];

            if (text[i] == pattern[j])
                ++j;

            if (j == pattern.Length)
                return i - j + 1;
        }

        return -1;
    }
}
