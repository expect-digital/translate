import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { TranslateService } from '@ngx-translate/core';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit, OnDestroy {
  readonly genderList = ['male', 'female', 'other'];

  readonly languageList = ['en', 'de'];

  readonly form = this.fb.nonNullable.group({
    count: this.fb.nonNullable.control(0),
    gender: this.fb.nonNullable.control(''),
  });

  readonly subscription = new Subscription();

  gender!: string;

  constructor(private fb: FormBuilder, private translate: TranslateService) {
    // this language will be used as a fallback when a translation isn't found in the current language
    translate.setDefaultLang('en');

    // the lang to use, if the lang isn't available, it will use the current loader to get them
    translate.use('en');
  }

  ngOnInit(): void {
    this.subscription.add(
      this.form.controls.gender.valueChanges.subscribe({
        next: (gender) => {
          this.translate.get('gender.options.' + gender).subscribe({
            next: (translate) => (this.gender = translate),
          });
        },
      })
    );
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }

  useLanguage(language: string): void {
    this.translate.use(language);
  }
}
