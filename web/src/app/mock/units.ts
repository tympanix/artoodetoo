import { Unit } from '../model';

export const UNITS = [
  { id: "example.EmailAction",
    description: "test1",
    input: [{name:"Receiver",type:"string"},{name:"Subject",type:"string"},{name:"Message",type:"string"}],
    output: []},
  { id:"example.StringConverter",
    description: "test2",
    input:[{name:"Format",type:"string"},{name:"Placeholder",type:"interface {}"}],
    output:[{name:"Formatted",type:"string"}]},
  { id:"example.PersonEvent",
    description: "test3",
    input:[],
    output:[{name:"Name",type:"string"},{name:"Age",type:"int"},{name:"Heigth",type:"float32"},{name:"Married",type:"bool"}]}
];
