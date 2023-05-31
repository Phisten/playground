public class Solution15
{
    public IList<IList<int>> ThreeSum(int[] nums)
    {
        var collect = new HashSet<string>();

        // 遞增排序
        var inputNums = new List<int>(nums);
        inputNums.Sort();
        // var len = inputNums.Count;


        ThreeSumSubFunc(collect, inputNums, 0, 0, 3,'', 3);

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

    public void ThreeSumSubFunc(
        HashSet<string> collect,
        List<int> orderNums,
        int target,
        int headIndex,
        int lastSum,
        string lastSumKey,
        int lastNums
    )
    {
        var len = orderNums.Count;

        // 與lastCount相加後正好使得數列總和為 target
        int exactlyNum = target - lastSum;

        // 數列拆分成三組： index:0 <= lesser < exactly < moreNum <= index:(orderNums.length-1)
        int lesserStart = headIndex;
        var exactlyNumStart = orderNums.FindIndex(lesserStart, x => exactlyNum == x);
        var moreNumStart = orderNums.FindIndex(exactlyNumStart + 1, x => exactlyNum != x);

        if (lastNums == 1)
        {
            collect.Add(lastSumKey + ',' + exactlyNum);
        }
    }

    // O(N^3) 超出時限
    public IList<IList<int>> Exhaustive(int[] nums)
    {
        var collect = new HashSet<string>();

        // 遞增排序
        var inputNums = new List<int>(nums);
        inputNums.Sort();
        var len = inputNums.Count;

        for (int j = 0; j < len - 2; j++)
        {
            for (int k = j + 1; k < len - 1; k++)
            {
                for (int l = k + 1; l < len; l++)
                {
                    if (inputNums[j] + inputNums[k] + inputNums[l] == 0)
                    {
                        var temp = inputNums[j] + "," + inputNums[k] + "," + inputNums[l];

                        collect.Add(temp);
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
