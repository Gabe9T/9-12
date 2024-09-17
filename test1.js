
class Histo {
    constructor(ages) {
      // Initialize 10 buckets with 0
      this.buckets = new Array(10).fill(0);
  
      // Fill the buckets based on the ages
      ages.forEach(age => {
        let bucketIndex = Math.floor(age / 10);
        this.buckets[bucketIndex] += 1;
      });
    }
  
    getHisto() {
      let total = this.buckets.reduce((acc, curr) => acc + curr, 0);
  
      // Normalize the buckets to percentages
      let normalizedBuckets = this.buckets.map(count => {
        let percentage = (count / total) * 100;
        return Math.round(percentage * 100) / 100;  // Truncate to 2 decimal places
      });
  
      return normalizedBuckets;
    }
  }
  
  // test 
  let histo = new Histo([1, 67, 99, 21, 55, 87, 23, 33, 11]);
  let result = histo.getHisto();
  console.log(result);
  
  