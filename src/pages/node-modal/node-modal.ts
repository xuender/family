import { Component } from '@angular/core';
import { NavParams, ViewController  } from 'ionic-angular';

import { Calendar } from 'ionic-native';
import { TreeNode } from '../../tree/tree-node';
import { NodeType } from '../../tree/node-type';

/**
 * 节点编辑页面
 */
@Component({
  selector: 'page-node-modal',
  templateUrl: 'node-modal.html'
})
export class NodeModal {
  node: TreeNode;
  // 可选择的父亲或母亲
  others: string[];
  // 父亲或母亲的称呼
  otherTitle: string;
  parentNode: TreeNode; // 父亲或母亲
  constructor(
    public params: NavParams,
    public viewCtrl: ViewController,
  ) {
    this.node = this.params.get('node');
    this.others = [];
    if (this.node.nt === NodeType.DEFAULT) { // 子女才需要设置父母
      const tree = this.params.get('tree');
      const old = this.params.get('old');
      if (tree) {
        this.setParent(tree.root, old);
        if (this.parentNode) {
          this.otherTitle = this.parentNode.gender ? '母亲' : '父亲';
          for (const c of this.parentNode.children) {
            if (c.nt !== NodeType.DEFAULT) {
              this.others.push(c.name);
            }
          }
        }
      }
    }
  }
  // 创建日历提醒
  createDob() {
    this.createCalendar(new Date(this.node.dob), `${this.node.name}生日`);
  }
  // 创建忌日提醒
  createDod() {
    this.createCalendar(new Date(this.node.dod), `${this.node.name}忌日`);
  }
  // 创建提醒
  createCalendar(time: Date, title: string) {
    const now = new Date();
    const start = new Date(now.getTime());
    start.setMonth(time.getMonth());
    start.setDate(time.getDate());
    if (start < now) {  // 已经过去，下年提醒
       start.setMonth(start.getMonth() + 12);
    }
    start.setHours(0);
    start.setMinutes(0);
    start.setSeconds(0);
    const end = new Date(start.getTime());
    end.setHours(24);
    end.setMinutes(0);
    end.setSeconds(0);
    const note = title + ' by 家谱';
    Calendar.createEventInteractivelyWithOptions(
      title, null, note, start, end, {recurrence: 'yearly', recurrenceInterval: 1}
    );
  }
  // 设置父节点
  setParent(root: TreeNode, node: TreeNode) {
    if (root.children) {
      for (const c of root.children) {
        if (c === node) {
          this.parentNode = root;
          return;
        }
        this.setParent(c, node);
      }
    }
  }
  ionViewWillEnter() {
    this.viewCtrl.setBackButtonText('返回');
  }
  /**
   * 取消
   */
  cancel() {
    console.debug('cancel');
    this.viewCtrl.dismiss();
  }
  /**
   * 确定
   */
  ok() {
    console.debug('ok');
    this.node.ua = new Date();
    this.viewCtrl.dismiss(this.node);
  }
}