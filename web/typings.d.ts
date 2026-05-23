import 'umi/typings';

declare module '*.less' {
  const classes: { [key: string]: string };
  export default classes;
}
