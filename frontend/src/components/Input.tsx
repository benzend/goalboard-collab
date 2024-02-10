import React, { ComponentProps } from 'react';
import cn from 'classnames';

export interface InputProps extends ComponentProps<'input'> {
  inputStyleType: InputStyleType;
}

export const Input: React.FunctionComponent<InputProps> = (props) => {
  const extended: ComponentProps<'input'> = {
    ...props,
    className: cn(getInputClasses(props.inputStyleType), props.className),
  };
  return <input {...extended} />;
};

export function getInputClasses(style: InputStyleType = 'default'): string {
  switch (style) {
    case 'w-full':
      return 'border border-black w-full px-2 h-10 rounded shadow-lg';
    case 'w-lg':
      return 'border border-black px-2 h-10 rounded shadow-lg w-20';
    case 'default':
    default:
      return 'border border-black px-2 h-10 rounded shadow-lg';
  }
}

export type InputStyleType = 'default' | 'w-full' | 'w-lg';
