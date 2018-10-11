/**
 * 這個node只是使用方式跟測試大量使用GRPC的情境下是否會卡住node.js
 * 測試下是不會卡住node的執行
*/
const caller = require('grpc-caller')
const path = require('path')
const PROTO_PATH = path.resolve(__dirname, './src/service/flopplayer2/flopplayer2.proto')
const client = caller('localhost:50051', PROTO_PATH, 'Flopplayer2')
const NS_PER_SEC = 1e9;
const MS_PER_NS = 1e-6
var gameCount = 0;

let inputCards = Array();
const pArray = [];
let inputCards1 = ['144', '131', '121', '122', '91', '24', '132'];
let diff
let time
let diff2
let time2
let stop = true;
let i = 0;
Promise.resolve(1)
.then(() => {
  time = process.hrtime();
  return client.CalculatorFlop({
    PublicPoker: [144, 131, 121],
    Data: [
      {Id: 'thai5', Pokers: [134, 132]},
      {Id: 'thai2', Pokers: [92, 123]}
    ]});
})
.then((data) => {
  diff = process.hrtime(time);
  console.log('data', data);
  console.log(`Benchmark took ${diff[0] * NS_PER_SEC + diff[1]} nanoseconds`);
  console.log(`Benchmark took ${(diff[0] * NS_PER_SEC + diff[1]) * MS_PER_NS} milliseconds`);

})
// client.PokerCalculator({ data: inputCards1 })

// function intervalFunc() {
//   ++i
//   console.log("Hello!!!!"+i);
// }
// Promise.resolve(1)
// .then(async () => {
//   const array = [];
//   console.log("being");
//   setInterval(intervalFunc,100);
//   time = process.hrtime();
//   for (let i = 0; i < 100000; i++) {
//     array.push(client.PokerCalculator({ data: inputCards1 }));
//   }
//   return Promise.all(array);
// })
// .then((data) => {
//   diff = process.hrtime(time);
  // console.log(`Benchmark took ${diff[0] * NS_PER_SEC + diff[1]} nanoseconds`);
  // console.log(`Benchmark took ${(diff[0] * NS_PER_SEC + diff[1]) * MS_PER_NS} milliseconds`);
//   stop = false;
//   clearInterval(intervalFunc);
// })
// .then(() => {
//   stop = false;
//   console.log(`stop1`);
// })
// Promise.resolve(1)
// .then(async () => {
//   const array = [];
//   console.log("being2");
//   time2 = process.hrtime();
//   for (let i = 0; i < 10000; i++) {
//     array.push(client.PokerCalculator({ data: inputCards1 }));
//   }
//   return Promise.all(array);
// })
// .then((data) => {
//   diff2 = process.hrtime(time2);
//   console.log(`Benchmark2 took ${diff2[0] * NS_PER_SEC + diff2[1]} nanoseconds`);
//   console.log(`Benchmark2 took ${(diff2[0] * NS_PER_SEC + diff2[1]) * MS_PER_NS} milliseconds`);
//   stop = false;
// })
// .then(() => {
//   stop = false;
//   console.log(`stop2`);
// })

// console.log(`Benchmark took ${diff[0] * NS_PER_SEC + diff[1]} nanoseconds`);
// console.log(`Benchmark took ${(diff[0] * NS_PER_SEC + diff[1]) * MS_PER_NS} milliseconds`);

