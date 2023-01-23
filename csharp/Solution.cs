public class Solution
{
    //leetCode 18: 4Sum
    public IList<IList<int>> FourSum(int[] nums, int target)
    {
        var collect = new List<IList<int>>();
        // opt.Add(new List<int>());
        System.Console.WriteLine("input nums length: " + nums.Length);

        // 遞增排序
        var inputNums = new List<int>(nums);
        inputNums.Sort();

        var test = FourSumSubFunc(collect, inputNums.ToArray(), target, 0, 0);
        System.Console.WriteLine(test);

        return collect;
    }

    static public List<int> FourSumSubFunc(
        IList<IList<int>> collect,
        int[] sortedNums,
        int target,
        int startIndex,
        // List<int> previousSelectNums,
        int lastSum
    )
    {
        // 終止條件
        if (startIndex >= sortedNums.Length)
        {
            return new List<int>();
        }
        else
        {
            for (int i = startIndex; i < sortedNums.Length; i++)
            {
                var currentNumber = sortedNums[i];
                var currentSum = lastSum + currentNumber;
                if (currentSum < target)
                {
                    var next = FourSumSubFunc(collect, target, i + 1, currentSum);
                }
                else if (currentSum == target)
                {
                    collect.Add(previousSelectNums.Add(currentNumber));
                }
            }
        }

        // var subIssue = FourSum(inputNums.ToArray(), )
    }
}
