public class Solution18
{
    //leetCode 18: 4Sum
    public IList<IList<int>> FourSum(int[] nums, int target)
    {
        var collect = new HashSet<string>();
        System.Console.WriteLine("input nums length: " + nums.Length);

        // 遞增排序
        var inputNums = new List<int>(nums);
        inputNums.Sort();
        var len = inputNums.Count;

        for (int i = 0; i < len - 3; i++)
        {
            for (int j = i + 1; j < len - 2; j++)
            {
                for (int k = j + 1; k < len - 1; k++)
                {
                    for (int l = k + 1; l < len; l++)
                    {
                        if (inputNums[i] + inputNums[j] + inputNums[k] + inputNums[l] == target)
                        {
                            var temp =
                                inputNums[i]
                                + ","
                                + inputNums[j]
                                + ","
                                + inputNums[k]
                                + ","
                                + inputNums[l];

                            collect.Add(temp);
                        }
                    }
                }
            }
        }

        var outputList = new List<IList<int>>();
        collect
            .ToList()
            .ForEach(x =>
            {
                var temp = x.Split(",").Select(y => int.Parse(y)).ToList();
                outputList.Add(temp);
            });

        return outputList;
    }
}
