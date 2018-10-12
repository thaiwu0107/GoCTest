/**
 * 這個node只是使用方式跟測試大量使用GRPC的情境下是否會卡住node.js
 * 測試下是不會卡住node的執行
*/
const caller = require('grpc-caller')
const path = require('path')
const PROTO_PATH1 = path.resolve(__dirname, './src/service/turn3/turn3.proto')
const PROTO_PATH2 = path.resolve(__dirname, './src/service/turn2/turn2.proto')
const PROTO_PATH3 = path.resolve(__dirname, './src/service/flopplayer2/flopplayer2.proto')
const PROTO_PATH4 = path.resolve(__dirname, './src/service/flopplayer3/flopplayer3.proto')
const client1 = caller('localhost:50051', PROTO_PATH1, 'Turn3')
const client2 = caller('localhost:50051', PROTO_PATH2, 'Turn2')
const client3 = caller('localhost:50051', PROTO_PATH3, 'Flopplayer2')
const client4 = caller('localhost:50051', PROTO_PATH4, 'Flopplayer3')
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
// turn3
.then(() => {
  time = process.hrtime();
  return client1.CalculatorFlop({
    PublicPoker: [144, 131, 121, 132],
    Data: [
      {Id: 'thai5', Pokers: [134, 133]},
      {Id: 'thai4', Pokers: [92, 123]},
      {Id: 'thai2', Pokers: [24, 22]}
    ]});
})
.then((data) => {
  diff = process.hrtime(time);
  console.log('turn3 => data', data);
  console.log(`turn3 => Benchmark took ${diff[0] * NS_PER_SEC + diff[1]} nanoseconds`);
  console.log(`turn3 => Benchmark took ${(diff[0] * NS_PER_SEC + diff[1]) * MS_PER_NS} milliseconds`);

})
// turn2
.then(() => {
  time = process.hrtime();
  return client2.CalculatorFlop({
    PublicPoker: [144, 131, 121, 132],
    Data: [
      {Id: 'thai5', Pokers: [134, 133]},
      {Id: 'thai4', Pokers: [92, 123]}
    ]});
})
.then((data) => {
  diff = process.hrtime(time);
  console.log('turn2 => data', data);
  console.log(`turn2 => Benchmark took ${diff[0] * NS_PER_SEC + diff[1]} nanoseconds`);
  console.log(`turn2 => Benchmark took ${(diff[0] * NS_PER_SEC + diff[1]) * MS_PER_NS} milliseconds`);

})
// flopplayer2
.then(() => {
  time = process.hrtime();
  return client3.CalculatorFlop({
    PublicPoker: [144, 131, 121],
    Data: [
      {Id: 'thai5', Pokers: [134, 133]},
      {Id: 'thai4', Pokers: [92, 123]}
    ]});
})
.then((data) => {
  diff = process.hrtime(time);
  console.log('flopplayer2 => data', data);
  console.log(`flopplayer2 => Benchmark took ${diff[0] * NS_PER_SEC + diff[1]} nanoseconds`);
  console.log(`flopplayer2 => Benchmark took ${(diff[0] * NS_PER_SEC + diff[1]) * MS_PER_NS} milliseconds`);

})
// flopplayer3
.then(() => {
  time = process.hrtime();
  return client4.CalculatorFlop({
    PublicPoker: [144, 131, 121],
    Data: [
      {Id: 'thai5', Pokers: [134, 133]},
      {Id: 'thai4', Pokers: [92, 123]},
      {Id: 'thai2', Pokers: [24, 22]}
    ]});
})
.then((data) => {
  diff = process.hrtime(time);
  console.log('flopplayer3 => data', data);
  console.log(`flopplayer3 => Benchmark took ${diff[0] * NS_PER_SEC + diff[1]} nanoseconds`);
  console.log(`flopplayer3 => Benchmark took ${(diff[0] * NS_PER_SEC + diff[1]) * MS_PER_NS} milliseconds`);

})
