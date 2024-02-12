import React, { ComponentProps } from 'react';
import cn from 'classnames';

export interface SelectProps extends ComponentProps<'select'> {
  selectStyleType: SelectStyleType;
}

export const Select: React.FunctionComponent<SelectProps> = (props) => {
  const extended: ComponentProps<'select'> = {
    ...props,
    className: cn(getSelectClasses(props.selectStyleType), props.className),
  };
  return <select {...extended}>{props.children}</select>;
};

export function getSelectClasses(style: SelectStyleType = 'default'): string {
  switch (style) {
    case 'w-full':
      return 'border border-black px-2 h-10 rounded shadow-lg bg-white w-full';
    case 'w-lg':
      return 'border border-black px-2 h-10 rounded shadow-lg bg-white w-20';
    case 'default':
    default:
      return 'border border-black px-2 h-10 rounded shadow-lg bg-white';
  }
}

export type SelectStyleType = 'default' | 'w-full' | 'w-lg';
