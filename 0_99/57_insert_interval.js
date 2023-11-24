var insert = function(intervals, newInterval) {
   if (intervals.length === 0) {
      return [newInterval];
   }

   let l = 0;
   let r = intervals.length - 1;

   const findInsertIndex = () => {
      while (l <= r) {
         let mid = Math.floor((l + r) / 2);
         if (intervals[mid][0] === newInterval[0]) {
            return mid;
         }

         if (intervals[mid][0] > newInterval[0]) {
            r = mid - 1;
         } else {
            l = mid + 1;
         }
      }

      return l;
   }

   const idx = findInsertIndex();

   const newIntervals = [];
   if (idx - 1 < 0) {
      newIntervals.push(newInterval);
   }
   for (let i = 0; i < intervals.length; i++) {
      newIntervals.push(intervals[i]);
      if (i === idx - 1) {
         newIntervals.push(newInterval);
      }
   }

   const ans = [];
   for (const [start, end] of newIntervals) {
      if (ans.length && start <= ans[ans.length - 1][1]) {
         ans[ans.length - 1][1] = Math.max(ans[ans.length - 1][1], end);
      } else {
         ans.push([start, end]);
      }
   }

   return ans;
};
